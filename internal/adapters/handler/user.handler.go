package handler

import (
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type userHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *userHandler {
	return &userHandler{userService}
}
