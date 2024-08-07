package controllers

import (
	"diving-log-book-service/internal/pkg/apihelper"
	"net/http"
)

type HealthcheckController struct{}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

func (h HealthcheckController) Alive(w http.ResponseWriter, r *http.Request) {
	apihelper.Response(w, "OK")
}
