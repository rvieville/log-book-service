package controllers

import (
	"diving-log-book-service/internal/pkg/apihelper"
	clogger "diving-log-book-service/internal/pkg/logger"
	"diving-log-book-service/internal/pkg/mux"
	"diving-log-book-service/internal/services"
	"diving-log-book-service/internal/types"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type DiveController struct {
	diveService *services.DiveService
}

func NewDiveController(diveService *services.DiveService) *DiveController {
	return &DiveController{
		diveService,
	}
}

func (fm DiveController) Create(w http.ResponseWriter, r *http.Request) {
	var body types.CreateDivePayload

	json.NewDecoder(r.Body).Decode(&body)

	validationErr := apihelper.Validate(body)
	if validationErr != nil {
		apihelper.ValidationError(w, validationErr.(validator.ValidationErrors))
		return
	}

	dive, err := fm.diveService.Create(&body)
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	apihelper.Response(w, dive)
}

func (fm DiveController) ReadAll(w http.ResponseWriter, r *http.Request) {
	logger := mux.GetLoggerFromContext(r.Context())

	logger.Info(clogger.LogMessage{
		Event: "dive_controller_readAll",
		Msg:   "Start fetching all dives",
	})

	dives, err := fm.diveService.ReadAll()
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	apihelper.Response(w, dives)
}
