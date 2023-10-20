package handler

import (
	"admin-catalogo-go/pkg/events"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type VideoRegisteredHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewVideoRegisteredHandler(rabbitMQChannel *amqp.Channel) *VideoRegisteredHandler {
	return &VideoRegisteredHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *VideoRegisteredHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Category created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"admin-catalogo", // exchange
		"register_video", // key name
		false,            // mandatory
		false,            // immediate
		msgRabbitmq,      // message to publish
	)
}
