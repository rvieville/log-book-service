package repositories_test

import (
	gormHelper "diving-log-book-service/internal/pkg/gorm"
	"diving-log-book-service/test/utils"
	"fmt"
	"testing"
)

func TestReadAllCountry(t *testing.T) {
	repo := utils.GetCountryRepo()
	filters := []gormHelper.Filter{}
	filters = append(filters, gormHelper.Filter{
		Expression: "id IN ?",
		Data: []any{
			[]any{1, 2, 3, 4},
		},
	})

	res, err := repo.ReadAll(filters)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
