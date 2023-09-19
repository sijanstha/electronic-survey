package domain

import (
	"time"
)

type PollState string
type Role string

const (
	PREPARED PollState = "PREPARED"
	STARTED  PollState = "STARTED"
	VOTING   PollState = "VOTING"
	FINISHED PollState = "FINISHED"
)

const (
	ROLE_ORGANIZER Role = "ROLE_ORGANIZER"
	ROLE_ADMIN     Role = "ROLE_ADMIN"
)

type BaseEntity struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	BaseEntity
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	HashPassword string `json:"hashPassword"`
	Role         Role   `json:"role"`
}

type Poll struct {
	BaseEntity
	Title       string    `json:"title"`
	Description string    `json:"description"`
	State       PollState `json:"state"`
	StartsAt    time.Time `json:"startsAt"`
	EndsAt      time.Time `json:"endsAt"`
	Timezone    string    `json:"timezone"`
}

type PollOrganizer struct {
	BaseEntity
	OrganizerId      int64 `json:"organizerId"`
	PollId           int64 `json:"pollId"`
	PrimaryOrganizer bool  `json:"primaryOrganizer"`
}

type PaginationDetails struct {
	Size       int `json:"pageSize"`
	Page       int `json:"page"`
	Total      int `json:"totalRecords"`
	TotalPages int `json:"totalPages"`
}

type PollWithOrganizerInfo struct {
	Poll
	FullName string `json:"primaryOrganizerName"`
	Email    string `json:"primaryOrganizerEmail"`
}

type PollPaginationDetails struct {
	PaginationDetails
	Data []PollWithOrganizerInfo `json:"data"`
}
