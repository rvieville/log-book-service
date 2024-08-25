package repositories_test

import (
	gormHelper "diving-log-book-service/internal/pkg/gorm"
	"diving-log-book-service/test/utils"
	"fmt"
	"testing"
)

func TestReadAllIsland(t *testing.T) {
	repo := utils.GetIslandRepo()
	filters := []gormHelper.Filter{}
	filters = append(filters, gormHelper.Filter{
		Expression: "slug like ?",
		Data: []any{
			fmt.Sprintf("%%%s%%", "cor"),
		},
	})

	res, err := repo.ReadAll(filters)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
