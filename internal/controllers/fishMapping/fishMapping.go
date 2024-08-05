package fishmapping

import (
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/repositories"
	fishmapping "diving-log-book-service/internal/services/fishMapping"
	"diving-log-book-service/internal/types"
	"encoding/json"
	"net/http"
)

type FishMapping struct {
	service *fishmapping.FishMapping
}

func New() *FishMapping {
	return &FishMapping{
		service: fishmapping.New(
			repositories.NewFishMappingRepository(db.DB),
		),
	}
}

func (fm FishMapping) Create(w http.ResponseWriter, r *http.Request) {
	var body types.CreateFishMappingPayload

	json.NewDecoder(r.Body).Decode(&body)

	validationErr := body.Validate()
	if validationErr != nil {
		apihelper.ValidationError(w, validationErr)
		return
	}

	fishMapping, err := fm.service.Create(body)
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	apihelper.Response(w, fishMapping)
}
