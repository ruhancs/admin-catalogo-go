package handler

import (
	"admin-catalogo-go/pkg/events"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type CategoryCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewCategoryCreatedHandler(rabbitMQChannel *amqp.Channel) *CategoryCreatedHandler {
	return &CategoryCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *CategoryCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, key string) {
	defer wg.Done()
	fmt.Printf("Category created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"admin-catalogo", // exchange
		key,           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}