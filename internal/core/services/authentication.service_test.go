package services

import (
	"fmt"
	"testing"
	"time"

	mockrepository "github.com/sijanstha/electronic-voting-system/internal/adapters/repository/mock"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	mockservice "github.com/sijanstha/electronic-voting-system/internal/core/services/mock"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func initTestForAuthentication(t *testing.T) (ports.AuthenticationService, *mockrepository.MockUserRepository, *mockservice.MockTokenService) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockrepository.NewMockUserRepository(ctrl)
	tokenService := mockservice.NewMockTokenService(ctrl)

	return NewAuthenticationService(userRepo, tokenService), userRepo, tokenService
}

func TestEmailValidationErrorForAuthentication(t *testing.T) {
	service, _, _ := initTestForAuthentication(t)

	resp, err := service.Authenticate(&domain.LoginRequest{
		Email:    "",
		Password: utils.RandomString(10),
	})

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "email cannot be null or empty")
}

func TestPasswordValidationErrorForAuthentication(t *testing.T) {
	service, _, _ := initTestForAuthentication(t)

	resp, err := service.Authenticate(&domain.LoginRequest{
		Email:    utils.RandomEmail(),
		Password: "",
	})

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "password cannot be null or empty")
}

func TestEmailNotExistsErrorForAuthentication(t *testing.T) {
	service, userRepo, _ := initTestForAuthentication(t)

	email := utils.RandomEmail()
	userRepo.EXPECT().FindByEmail(email).Times(1).Return(nil, &commonError.ErrNotFound{Message: fmt.Sprintf("%s not found", email)})

	resp, err := service.Authenticate(&domain.LoginRequest{
		Email:    email,
		Password: utils.RandomString(10),
	})

	var notFoundErr *commonError.ErrNotFound
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &notFoundErr)
	require.ErrorContains(t, err, "not found")
}

func TestInvalidPasswordErrorForAuthentication(t *testing.T) {
	service, userRepo, _ := initTestForAuthentication(t)

	user := randomUser(utils.RandomString(10))
	userRepo.EXPECT().FindByEmail(user.Email).Times(1).Return(user, nil)

	resp, err := service.Authenticate(&domain.LoginRequest{
		Email:    user.Email,
		Password: utils.RandomString(10),
	})

	var notAuthorizedErr *commonError.ErrUnauthorized
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &notAuthorizedErr)
	require.ErrorContains(t, err, "invalid email or password")
}

func TestForAuthentication(t *testing.T) {
	service, userRepo, tokenService := initTestForAuthentication(t)

	password := utils.RandomString(10)
	user := randomUser(password)
	token := utils.RandomString(100)

	tokenService.EXPECT().Generate(*user).Times(1).Return(token, nil)
	userRepo.EXPECT().FindByEmail(user.Email).Times(1).Return(user, nil)

	resp, err := service.Authenticate(&domain.LoginRequest{
		Email:    user.Email,
		Password: password,
	})

	require.NotNil(t, resp)
	require.Nil(t, err)
	require.EqualValues(t, &domain.LoginResponse{Token: token}, resp)
}

func randomUser(password string) *domain.User {
	return &domain.User{
		FirstName:    utils.RandomFirstName(),
		LastName:     utils.RandomLastName(),
		Email:        utils.RandomEmail(),
		HashPassword: utils.HashPassword(password),
		Role:         domain.Role(utils.RandomRole()),
		BaseEntity: domain.BaseEntity{
			Id:        utils.RandomInt(10, 100),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
