package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

type authenticationHandler struct {
	userService ports.UserService
	authService ports.AuthenticationService
}

func NewAuthenticationHandler(userService ports.UserService, authService ports.AuthenticationService) *authenticationHandler {
	return &authenticationHandler{userService, authService}
}

func (h *authenticationHandler) HandleUserAuthentication(w http.ResponseWriter, r *http.Request) error {
	loginRequest := new(domain.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
		return err
	}

	resp, err := h.authService.Authenticate(loginRequest)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, domain.NewApiResponse(resp))
}

func (h *authenticationHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) error {
	createUserReq := new(domain.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(createUserReq); err != nil {
		return err
	}

	res, err := h.userService.SaveUser(createUserReq)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusAccepted, domain.NewApiResponse(res))
}
