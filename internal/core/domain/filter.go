package domain

type PaginationFilter struct {
	Limit  int64  `json:"limit" default:"10"`
	Page   int64  `json:"page" default:"1"`
	Sort   string `json:"sort" default:"asc"`
	SortBy string `json:"sortBy" default:"updated_at"`
}

type PollFilter struct {
	Id    int64
	Title string
}

type PollListFilter struct {
	PaginationFilter
	States              []PollState `json:"state"`
	OrganizerId        int64     `json:"organizerId"`
	FilterPrimaryOwner *bool
}
