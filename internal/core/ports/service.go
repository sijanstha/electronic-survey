package ports

import (
	"context"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
)

type PollService interface {
	SavePoll(context.Context, *domain.CreatePollRequest) (*domain.Poll, error)
	GetPollById(int64) (*domain.PollInfo, error)
	GetAllPoll(domain.PollListFilter) (*domain.PollPaginationDetails, error)
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
