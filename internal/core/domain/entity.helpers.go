package domain

import (
	"time"

	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

func NewPoll(id int64, title, description string, startsAt, endsAt time.Time, tz string, state PollState, forUpdate bool) *Poll {
	return &Poll{
		Title:       title,
		Description: description,
		StartsAt:    startsAt,
		EndsAt:      endsAt,
		Timezone:    tz,
		State:       state,
		BaseEntity:  newBaseEntity(id, forUpdate),
	}
}

func NewUser(request CreateUserRequest) *User {
	return &User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        request.Email,
		HashPassword: utils.HashPassword(request.Password),
		Role:         ROLE_ORGANIZER,
		BaseEntity:   newBaseEntity(0, false),
	}
}

func NewPollOrganizer(organizerId, pollId int64, isPrimaryOrganizer bool) *PollOrganizer {
	return &PollOrganizer{
		OrganizerId:      organizerId,
		PollId:           pollId,
		PrimaryOrganizer: isPrimaryOrganizer,
		BaseEntity:       newBaseEntity(0, false),
	}
}

func NewParticipantList(id int64, name string, emails []string, forUpdate bool) *ParticipantList {
	return &ParticipantList{
		Name:       name,
		Emails:     emails,
		BaseEntity: newBaseEntity(id, forUpdate),
	}
}

func newBaseEntity(id int64, forUpdate bool) BaseEntity {
	base := BaseEntity{}
	if !forUpdate {
		base.CreatedAt = time.Now().UTC()
	}
	base.Id = id
	base.UpdatedAt = time.Now().UTC()
	return base
}
