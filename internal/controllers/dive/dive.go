package dive

import (
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/services/dive"
	"diving-log-book-service/internal/types"
	"encoding/json"
	"net/http"
)

type Dive struct {
	service *dive.Dive
}

func New() *Dive {
	return &Dive{
		service: dive.New(
			repositories.NewDiveRepository(db.DB),
		),
	}
}

func (fm Dive) Create(w http.ResponseWriter, r *http.Request) {
	var body types.CreateDivePayload

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
