package services

import (
	"context"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
)

func getCurrentLoggedInUser(ctx context.Context) (*domain.Claims, error) {
	principal := ctx.Value("principal")
	if principal == nil {
		return nil, &commonError.ErrUnauthorized{Message: "unauthorized"}
	}

	claims, ok := principal.(*domain.Claims)
	if !ok {
		return nil, &commonError.ErrInternalServer{Message: "couldn't parse jwt claims"}
	}

	return claims, nil
}
