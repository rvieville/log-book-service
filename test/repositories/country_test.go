package repositories_test

import (
	gormHelper "diving-log-book-service/internal/pkg/gorm"
	"diving-log-book-service/test/utils"
	"fmt"
	"testing"
)

func TestReadAll(t *testing.T) {
	repo := utils.GetCountryRepo()
	filters := []gormHelper.Filter{}
	filters = append(filters, gormHelper.Filter{
		Expression: "slug like ?",
		Data: []any{
			fmt.Sprintf("%%%s%%", "mala"),
		},
	})

	res, err := repo.ReadAll(filters)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
