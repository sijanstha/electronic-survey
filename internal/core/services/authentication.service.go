package services

import (
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

type authenticationService struct {
	userRepo   ports.UserRepository
	jwtService ports.TokenService
}

func NewAuthenticationService(userRepo ports.UserRepository, jwtService ports.TokenService) ports.AuthenticationService {
	return &authenticationService{userRepo, jwtService}
}

func (s *authenticationService) Authenticate(req *domain.LoginRequest) (*domain.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(req.Password, user.HashPassword) {
		return nil, &commonError.ErrUnauthorized{Message: "invalid email or password"}
	}

	token, err := s.jwtService.Generate(*user)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{Token: token}, nil
}
