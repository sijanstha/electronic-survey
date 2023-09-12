package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type jwtService struct {
}

func NewJwtService() ports.TokenService {
	return &jwtService{}
}

func (s *jwtService) Generate(user domain.User) (string, error) {
	expirationTimeInMinutes := os.Getenv("JWT_EXPIRATION_TIME")
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	timeN, err := strconv.Atoi(expirationTimeInMinutes)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(time.Duration(timeN) * time.Minute)
	claims := &domain.Claims{
		Email: user.Email,
		Id:    user.Id,
		Role:  string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "evs@evs.com",
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ID:        fmt.Sprintf("%d", user.Id),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}

func (s *jwtService) Validate(token string) (*domain.Claims, error) {
	claims := &domain.Claims{}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, &commonError.ErrUnauthorized{Message: "invalid token"}
		}

		return nil, err
	}

	return claims, nil
}
