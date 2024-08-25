package controllers

import (
	"diving-log-book-service/internal/pkg/apihelper"
	gormHelper "diving-log-book-service/internal/pkg/gorm"
	"diving-log-book-service/internal/services"
	"net/http"
	"strings"
)

type IslandController struct {
	islandService *services.IslandService
}

func NewIslandContorller(islandService *services.IslandService) *IslandController {
	return &IslandController{
		islandService,
	}
}

func (i IslandController) ReadAll(w http.ResponseWriter, r *http.Request) {
	var filters []gormHelper.Filter
	query := r.URL.Query()
	ids := query.Get("ids")
	countryID := query.Get("countryIds")

	if ids != "" {
		idsArray := strings.Split(ids, ",")
		filters = append(filters, gormHelper.Filter{
			Expression: "id IN ?",
			Data: []any{
				idsArray,
			},
		})
	}

	if countryID != "" {
		countryIds := strings.Split(countryID, ",")
		filters = append(filters, gormHelper.Filter{
			Expression: "country_id IN ?",
			Data: []any{
				countryIds,
			},
		})
	}

	countries, err := i.islandService.ReadAll(filters)
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	apihelper.Response(w, countries)
}
