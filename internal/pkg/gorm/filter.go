package gormHelper

import "gorm.io/gorm"

type Filter struct {
	Expression string
	Data       []any
}

func CreateFilter(filters []Filter) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, v := range filters {
			db.Where(v.Expression, v.Data...)
		}

		return db
	}
}
