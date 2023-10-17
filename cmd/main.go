package main

import (
	"admin-catalogo-go/internal/application/factory"
	"admin-catalogo-go/internal/event/handler"
	"admin-catalogo-go/internal/infra/web"
	"admin-catalogo-go/pkg/events"
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db,err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("CategoryCreated", &handler.CategoryCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("CategoryDeleted", &handler.CategoryDeletedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createCategoryUseCase := factory.CreateCategoryUseCaseFactory(db,eventDispatcher)
	listCategoryUseCase := factory.ListCategoryUsecaseFactory(db)
	getCategoryUseCase := factory.GetCategoryByIDUsecaseFactory(db)
	deleteCategoryUseCase := factory.DeleteCategoryUsecaseFactory(db,eventDispatcher)

	app := web.NewApplication(*createCategoryUseCase, *getCategoryUseCase, *deleteCategoryUseCase,*listCategoryUseCase)

	app.Server()
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONN"))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}