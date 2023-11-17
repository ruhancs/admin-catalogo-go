package handler

import (
	"admin-catalogo-go/pkg/events"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type VideoFileUploadedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewVideoFileUploadedHandler(rabbitMQChannel *amqp.Channel) *VideoPublishedHandler {
	return &VideoPublishedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *VideoFileUploadedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Video To Proccess Sending: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"admin-catalogo", // exchange
		"video-proccess", // key name
		false,            // mandatory
		false,            // immediate
		msgRabbitmq,      // message to publish
	)
}
