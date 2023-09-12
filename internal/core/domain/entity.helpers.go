package domain

import (
	"time"

	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

func NewPoll(title, description string, startsAt, endsAt time.Time) *Poll {
	return &Poll{
		Title:       title,
		Description: description,
		StartsAt:    startsAt,
		EndsAt:      endsAt,
		State:       PREPARED,
		BaseEntity:  newBaseEntity(false),
	}
}

func NewUser(request CreateUserRequest) *User {
	return &User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        request.Email,
		HashPassword: utils.HashPassword(request.Password),
		Role:         ROLE_ORGANIZER,
		BaseEntity:   newBaseEntity(false),
	}
}

func NewPollOrganizer(organizerId, pollId int64, isPrimaryOrganizer bool) *PollOrganizer {
	return &PollOrganizer{
		OrganizerId:      organizerId,
		PollId:           pollId,
		PrimaryOrganizer: isPrimaryOrganizer,
		BaseEntity:       newBaseEntity(false),
	}
}

func newBaseEntity(forUpdate bool) BaseEntity {
	base := BaseEntity{}
	if !forUpdate {
		base.CreatedAt = time.Now().UTC()
	}
	base.UpdatedAt = time.Now().UTC()
	return base
}
