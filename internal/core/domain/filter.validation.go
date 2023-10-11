package domain

import (
	"reflect"

	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

// TODO: Make map of sorting field based upon entity, Currently this list contains generic sorting fields for all the entities
var (
	validSortingField = map[string]bool{
		"title":       true,
		"state":       true,
		"created_at":  true,
		"updated_at":  true,
		"description": true,
		"id":          true,
	}

	validSortingOrder = map[string]bool{
		"asc":  true,
		"desc": true,
	}
)

func (filter *PaginationFilter) Validate() error {
	typ := reflect.TypeOf(*filter)
	if filter.Limit == 0 {
		field, _ := typ.FieldByName("Limit")
		filter.Limit = utils.ParseInteger(field.Tag.Get("default"))
	}

	if filter.Page == 0 {
		field, _ := typ.FieldByName("Page")
		filter.Page = utils.ParseInteger(field.Tag.Get("default"))
	}

	if filter.Sort == "" || len(filter.Sort) == 0 {
		field, _ := typ.FieldByName("Sort")
		filter.Sort = field.Tag.Get("default")
	}

	if filter.SortBy == "" || len(filter.SortBy) == 0 {
		field, _ := typ.FieldByName("SortBy")
		filter.SortBy = field.Tag.Get("default")
	}

	if filter.SortBy != "" && len(filter.SortBy) > 0 {
		if !validSortingField[filter.SortBy] {
			return commonError.ErrInvalidSortingField
		}
	}
	if filter.Sort != "" && len(filter.Sort) > 0 {
		if !validSortingOrder[filter.Sort] {
			return commonError.ErrInvalidSortingOrder
		}
	}

	return nil
}
