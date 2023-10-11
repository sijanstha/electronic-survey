package domain

import (
	"net/http"

	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

type PaginationFilter struct {
	Limit           int64  `json:"limit" default:"10"`
	Page            int64  `json:"page" default:"1"`
	Sort            string `json:"sort" default:"asc"`
	SortBy          string `json:"sortBy" default:"updated_at"`
	SearchParameter string `json:"search"`
}

type PollFilter struct {
	Id    int64
	Title string
}

type PollListFilter struct {
	PaginationFilter
	States             []PollState `json:"state"`
	OrganizerId        int64       `json:"organizerId"`
	FilterPrimaryOwner *bool
}

type ParticipantListFilter struct {
	PaginationFilter
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	OrganizerId int64  `json:"organizerId"`
}

func ParsePaginationRequest(r *http.Request) PaginationFilter {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")
	sortBy := r.URL.Query().Get("sortBy")
	searchParameter := r.URL.Query().Get("search")

	return PaginationFilter{
		Limit:           utils.ParseInteger(limit),
		Page:            utils.ParseInteger(page),
		Sort:            sort,
		SortBy:          sortBy,
		SearchParameter: searchParameter,
	}
}
