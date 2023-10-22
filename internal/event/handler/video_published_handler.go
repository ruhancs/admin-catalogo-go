package handler

import (
	"admin-catalogo-go/pkg/events"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type VideoPublishedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewVideoPublishedHandler(rabbitMQChannel *amqp.Channel) *VideoPublishedHandler {
	return &VideoPublishedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *VideoPublishedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Video Published: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"admin-catalogo", // exchange
		"publish_video",  // key name
		false,            // mandatory
		false,            // immediate
		msgRabbitmq,      // message to publish
	)
}
