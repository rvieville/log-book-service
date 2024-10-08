package controllers

import (
	"diving-log-book-service/internal/pkg/apihelper"
	gormHelper "diving-log-book-service/internal/pkg/gorm"
	"diving-log-book-service/internal/services"
	"net/http"
	"strings"
)

type CountryController struct {
	countryService *services.CountryService
}

func NewCountryController(countryService *services.CountryService) *CountryController {
	return &CountryController{
		countryService,
	}
}

func (c CountryController) ReadAll(w http.ResponseWriter, r *http.Request) {
	var filters []gormHelper.Filter
	query := r.URL.Query()
	ids := query.Get("ids")

	if ids != "" {
		idsArray := strings.Split(ids, ",")
		filters = append(filters, gormHelper.Filter{
			Expression: "id IN ?",
			Data: []any{
				idsArray,
			},
		})
	}

	countries, err := c.countryService.ReadAll(filters)
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	apihelper.Response(w, countries)
}
