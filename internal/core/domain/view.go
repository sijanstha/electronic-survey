package domain

import "github.com/sijanstha/electronic-voting-system/internal/core/utils"

type PollOwnerDatabaseView struct {
	Id           int64  `json:"id"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	PrimaryOwner int    `json:"primaryOwner"`
}

type PollDatabaseView struct {
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	State       PollState               `json:"state"`
	StartsAt    string                  `json:"startsAt"`
	EndsAt      string                  `json:"endsAt"`
	Id          int64                   `json:"id"`
	CreatedAt   string                  `json:"createdAt"`
	UpdatedAt   string                  `json:"updatedAt"`
	PollOwners  []PollOwnerDatabaseView `json:"pollOwners"`
}

type PollOwnerInfo struct {
	Id           int64  `json:"id"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	PrimaryOwner bool   `json:"primaryOwner"`
}

type PollInfo struct {
	Poll
	PollOwners []PollOwnerInfo `json:"pollOwners"`
}

func (p *PollOwnerDatabaseView) ToPollOwnerInfo() PollOwnerInfo {
	isPrimaryOwner := false
	if p.PrimaryOwner == 1 {
		isPrimaryOwner = true
	}

	return PollOwnerInfo{
		Id:           p.Id,
		FullName:     p.FullName,
		Email:        p.Email,
		PrimaryOwner: isPrimaryOwner,
	}
}

func (p *PollDatabaseView) ToPollInfo() *PollInfo {
	pollOwners := make([]PollOwnerInfo, 0)
	for _, data := range p.PollOwners {
		pollOwners = append(pollOwners, data.ToPollOwnerInfo())
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
		},
		PollOwners: pollOwners,
	}

	return pollInfo
}
