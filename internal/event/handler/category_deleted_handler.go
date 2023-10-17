package handler

import (
	"admin-catalogo-go/pkg/events"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type CategoryDeletedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewCategoryDeletedHandler(rabbitMQChannel *amqp.Channel) *CategoryDeletedHandler {
	return &CategoryDeletedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *CategoryDeletedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Category Deleted: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"admin-catalogo",  // exchange
		"delete_category", // key name
		false,             // mandatory
		false,             // immediate
		msgRabbitmq,       // message to publish
	)
}
