package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

type ApiError struct {
	Error     string `json:"error"`
	Timestamp string `json:"timestmap"`
}

func NewApiError(error string) *ApiError {
	return &ApiError{Error: error, Timestamp: utils.Now()}
}

type ApiResponse struct {
	Body      any    `json:"body"`
	Timestamp string `json:"timestmap"`
}

func NewApiResponse(body any) *ApiResponse {
	return &ApiResponse{Body: body, Timestamp: utils.Now()}
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
