package routes

import (
	"diving-log-book-service/internal/controllers"
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func DiveRouteInit(router *mux.Router) {
	diveRepo := repositories.NewDiveRepository(db.DB)
	diveFishRepo := repositories.NewDiveFishRepository(db.DB)
	mediaRepo := repositories.NewMediaRepository(db.DB)
	diveFishService := services.NewDiveFishService(diveFishRepo)
	mediaService := services.NewMediaService(mediaRepo)
	diveService := services.NewDiveService(&services.DiveServiceConfig{
		DiveRepo:        diveRepo,
		DiveFishService: diveFishService,
		MediaService:    mediaService,
	})
	controller := controllers.NewDiveController(diveService)

	group := router.PathPrefix("/dive").Subrouter()

	group.HandleFunc("/create", controller.Create).Methods(http.MethodPost)
	group.HandleFunc("/list", controller.ReadAll).Methods(http.MethodGet)
}
