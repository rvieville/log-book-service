package routes

import (
	"diving-log-book-service/internal/controllers"
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/middlewares"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func DiveRouteInit(router *mux.Router) {
	diveRepo := repositories.NewDiveRepository(db.DB)
	diveFishRepo := repositories.NewDiveFishRepository(db.DB)
	diveFishService := services.NewDiveFishService(diveFishRepo)
	diveService := services.NewDiveService(&services.DiveServiceConfig{
		DiveRepo:        diveRepo,
		DiveFishService: diveFishService,
	})
	controller := controllers.NewDiveController(diveService)

	group := router.PathPrefix("/dive").Subrouter()
	group.Use(middlewares.LoggingMiddleware)

	group.HandleFunc("/create", controller.Create).Methods(http.MethodPost)
	group.HandleFunc("/list", controller.ReadAll).Methods(http.MethodGet)
}
