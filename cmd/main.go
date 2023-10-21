package main

import (
	"admin-catalogo-go/internal/application/factory"
	"admin-catalogo-go/internal/event/handler"
	"admin-catalogo-go/internal/infra/web"
	"admin-catalogo-go/pkg/events"
	"database/sql"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

var (
	s3Client *s3.S3
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("REGION")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("PK"),
				os.Getenv("SK"),
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
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
	eventDispatcher.Register("VideoRegistered", &handler.VideoRegisteredHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createCategoryUseCase := factory.CreateCategoryUseCaseFactory(db, eventDispatcher)
	listCategoryUseCase := factory.ListCategoryUsecaseFactory(db)
	getCategoryUseCase := factory.GetCategoryByIDUsecaseFactory(db)
	deleteCategoryUseCase := factory.DeleteCategoryUsecaseFactory(db, eventDispatcher)

	categoryIDValidator := factory.CategoryIDValidator(db)
	registerVideoUseCase := factory.RegisterVideoUseCaseFactory(db, eventDispatcher, s3Client, categoryIDValidator)

	app := web.NewApplication(
		*createCategoryUseCase,
		*getCategoryUseCase,
		*deleteCategoryUseCase,
		*listCategoryUseCase,
		*registerVideoUseCase,
	)

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
