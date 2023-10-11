package ports

import (
	"context"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
)

type PollService interface {
	SavePoll(context.Context, *domain.CreatePollRequest) (*domain.Poll, error)
	UpdatePoll(context.Context, *domain.UpdatePollRequest) (*domain.Poll, error)
	GetPollById(int64) (*domain.PollInfo, error)
	GetAllPoll(context.Context, domain.PollListFilter) (*domain.PollPaginationDetails, error)
}

type UserService interface {
	SaveUser(*domain.CreateUserRequest) (*domain.User, error)
}

type AuthenticationService interface {
	Authenticate(*domain.LoginRequest) (*domain.LoginResponse, error)
}

type TokenService interface {
	Generate(domain.User) (string, error)
	Validate(token string) (*domain.Claims, error)
}

type ParticipantListService interface {
	SaveParticipantList(context.Context, *domain.CreateParticipantListRequest) (*domain.ParticipantList, error)
	UpdateParticipantList(context.Context, *domain.UpdateParticipantList) (*domain.ParticipantList, error)
	GetParticipantListById(context.Context, int64) (*domain.ParticipantList, error)
	GetParticipantListByName(context.Context, string) (*domain.ParticipantList, error)
	GetParticipantList(context.Context, *domain.ParticipantListFilter) (*domain.ParticipantPaginationDetails, error)
}
