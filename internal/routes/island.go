package routes

import (
	"diving-log-book-service/internal/controllers"
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/services"

	"github.com/gorilla/mux"
)

func IslandRouteInit(router *mux.Router) {
	islandRepo := repositories.NewIslandRepository(db.DB)
	islandService := services.NewIslandService(islandRepo)
	islandController := controllers.NewIslandContorller(islandService)

	group := router.PathPrefix("/island").Subrouter()

	group.HandleFunc("/list", islandController.ReadAll)
}
