package ports

import (
	"time"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
)

type PollRepository interface {
	SavePoll(*domain.Poll) (*domain.Poll, error)
	UpdatePoll(*domain.Poll) (*domain.Poll, error)
	FindPoll(domain.PollFilter) (*domain.PollInfo, error)
	FindAllPoll(domain.PollListFilter) (*domain.PollPaginationDetails, error)
	FindAllPollInStartedStateInDateRange(from time.Time, to time.Time) ([]*domain.Poll, error)
	FindAllPollInVotingStateInDateRange(from time.Time, to time.Time) ([]*domain.Poll, error)
}

type UserRepository interface {
	SaveUser(*domain.User) (*domain.User, error)
	FindByEmail(string) (*domain.User, error)
}

type PollOrganizerRepository interface {
	SavePollOrganizer(*domain.PollOrganizer) (*domain.PollOrganizer, error)
}

type ParticipantListRepository interface {
	SaveParticipantList(*domain.ParticipantList) (*domain.ParticipantList, error)
	UpdateParticipantList(*domain.ParticipantList) (*domain.ParticipantList, error)
	FindParticipantList(*domain.ParticipantListFilter) (*domain.ParticipantList, error)
	FindAllParticipantList(*domain.ParticipantListFilter) (*domain.ParticipantPaginationDetails, error)
}