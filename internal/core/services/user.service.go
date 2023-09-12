package services

import (
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) SaveUser(req *domain.CreateUserRequest) (*domain.User, error) {
	user, err := req.Validate()
	if err != nil {
		return nil, err
	}

	user, err = s.userRepo.SaveUser(user)
	if err != nil {
		return nil, err
	}

	user.HashPassword = ""
	return user, nil
}
