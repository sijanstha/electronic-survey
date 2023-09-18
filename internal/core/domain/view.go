package domain

import "github.com/sijanstha/electronic-voting-system/internal/core/utils"

type PollOrganizerDatabaseView struct {
	Id               int64  `json:"id"`
	FullName         string `json:"fullName"`
	Email            string `json:"email"`
	PrimaryOrganizer int    `json:"primaryOrganizer"`
}

type PollDatabaseView struct {
	Title          string                      `json:"title"`
	Description    string                      `json:"description"`
	State          PollState                   `json:"state"`
	Timezone       string                      `json:"timezone"`
	StartsAt       string                      `json:"startsAt"`
	EndsAt         string                      `json:"endsAt"`
	Id             int64                       `json:"id"`
	CreatedAt      string                      `json:"createdAt"`
	UpdatedAt      string                      `json:"updatedAt"`
	PollOrganizers []PollOrganizerDatabaseView `json:"pollOrganizers"`
}

type PollOrganizerInfo struct {
	Id               int64  `json:"id"`
	FullName         string `json:"fullName"`
	Email            string `json:"email"`
	PrimaryOrganizer bool   `json:"primaryOrganizer"`
}

type PollInfo struct {
	Poll
	PollOrganizers []PollOrganizerInfo `json:"pollOrganizers"`
}

func (p *PollOrganizerDatabaseView) ToPollOrganizerInfo() PollOrganizerInfo {
	isPrimaryOrganizer := false
	if p.PrimaryOrganizer == 1 {
		isPrimaryOrganizer = true
	}

	return PollOrganizerInfo{
		Id:               p.Id,
		FullName:         p.FullName,
		Email:            p.Email,
		PrimaryOrganizer: isPrimaryOrganizer,
	}
}

func (p *PollDatabaseView) ToPollInfo() *PollInfo {
	pollOrganizers := make([]PollOrganizerInfo, 0)
	for _, data := range p.PollOrganizers {
		pollOrganizers = append(pollOrganizers, data.ToPollOrganizerInfo())
	}

	pollInfo := &PollInfo{
		Poll: Poll{
			BaseEntity: BaseEntity{
				Id:        p.Id,
				CreatedAt: utils.FormatDateTime(p.CreatedAt),
				UpdatedAt: utils.FormatDateTime(p.UpdatedAt),
			},
			Title:       p.Title,
			Description: p.Description,
			StartsAt:    utils.FormatDateTime(p.StartsAt),
			EndsAt:      utils.FormatDateTime(p.EndsAt),
			State:       p.State,
			Timezone:    p.Timezone,
		},
		PollOrganizers: pollOrganizers,
	}

	return pollInfo
}
