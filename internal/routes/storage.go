package routes

import (
	"diving-log-book-service/internal/controllers"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/services"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/gorilla/mux"
)

func StorageRouteInit(router *mux.Router) {
	config := &aws.Config{
		Region:   aws.String(os.Getenv("STORAGE_REGION")),
		Endpoint: aws.String(os.Getenv("STORAGE_URL")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("STORAGE_ACCESS_KEY"),
			os.Getenv("STORAGE_SECRET_KEY"),
			"", // a token will be created when the session it's used.
		),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
	}
	storageRepo := repositories.NewStorageRepository(config)
	storageService := services.NewStorageService(storageRepo)
	controller := controllers.NewStorageController(storageService)

	group := router.PathPrefix("/storage").Subrouter()

	group.HandleFunc("", controller.Get).Methods(http.MethodGet)
	group.HandleFunc("/upload", controller.Upload).Methods(http.MethodPost)
}
