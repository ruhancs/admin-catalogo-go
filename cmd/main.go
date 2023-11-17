package main

import (
	"admin-catalogo-go/internal/application/factory"
	"admin-catalogo-go/internal/event/handler"
	"admin-catalogo-go/internal/infra/web"
	"admin-catalogo-go/pkg/events"
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var (
	s3Client *s3.S3
	totalErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "admin_catalogo_errors",
		Help: "Total errors",
	})
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

	prometheus.MustRegister(totalErrors)
}


func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//logConfiguration := zap.Config{
	//	OutputPaths: []string{"stdout"},
	//	Level: zap.NewAtomicLevelAt(zapcore.InfoLevel),
	//	Encoding: "json",
	//	EncoderConfig: zapcore.EncoderConfig{
	//		MessageKey: "Msg",
	//		LevelKey: "Level",
	//		TimeKey: "Time",
	//		NameKey: "Name",
	//		EncodeTime: zapcore.ISO8601TimeEncoder,
	//	},
	//}
	//myLogger,_ := logConfiguration.Build()

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		logger.Error("failed to connet to db",
		zap.String("DB Connection", "failed to connect"),
		)
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel(sugar)

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
	eventDispatcher.Register("VideoFileUploaded", &handler.VideoFileUploadedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("VideoPublished", &handler.VideoPublishedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createCategoryUseCase := factory.CreateCategoryUseCaseFactory(db, eventDispatcher)
	listCategoryUseCase := factory.ListCategoryUsecaseFactory(db)
	getCategoryUseCase := factory.GetCategoryByIDUsecaseFactory(db)
	deleteCategoryUseCase := factory.DeleteCategoryUsecaseFactory(db, eventDispatcher)

	categoryIDValidator := factory.CategoryIDValidator(db)
	registerVideoFilesUseCase := factory.RegisterVideoFileUseCaseFactory(db, eventDispatcher, s3Client, categoryIDValidator)
	registerVideoMetaUseCase := factory.RegisterVideoMetaUseCaseFactory(db, eventDispatcher, s3Client, categoryIDValidator)
	listVideosUseCase := factory.ListVideosUsecaseFactory(db)
	getVideoByIDUseCase := factory.GetVideoByIDUsecaseFactory(db)
	getVideoByCategoryUseCase := factory.GetVideoByCategoryUsecaseFactory(db)
	updateVideoToPublishUseCase := factory.UpdateVideoPublishedUseCaseFactory(db, eventDispatcher)

	app := web.NewApplication(
		totalErrors,
		*createCategoryUseCase,
		*getCategoryUseCase,
		*deleteCategoryUseCase,
		*listCategoryUseCase,
		*registerVideoFilesUseCase,
		*registerVideoMetaUseCase,
		*listVideosUseCase,
		*getVideoByIDUseCase,
		*getVideoByCategoryUseCase,
		*updateVideoToPublishUseCase,
	)

	http.Handle("/metrics", promhttp.Handler())
	go func ()  {
		http.ListenAndServe(":8080", nil)
	}()
	
	app.Server()

}

func getRabbitMQChannel(sugar *zap.SugaredLogger) *amqp.Channel {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONN"))
	sugar.Infow("connecting to rabbitmq",
		"attempt", 3,
		"backoff", time.Second,
	)
	if err != nil {
		sugar.Errorw("Error to connect rabbitmq",
			"backoff", time.Second,
		)
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		sugar.Error("Error to connect on rabbitmq channel",
			"backoff", time.Second,
		)
		panic(err)
	}
	return ch
}
