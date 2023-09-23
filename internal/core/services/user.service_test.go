package services

import (
	"testing"

	mockrepository "github.com/sijanstha/electronic-voting-system/internal/adapters/repository/mock"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func initTestForUserService(t *testing.T) (ports.UserService, *mockrepository.MockUserRepository) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockrepository.NewMockUserRepository(ctrl)
	return NewUserService(userRepo), userRepo
}

func TestFirstNameValidationErrorForRegistration(t *testing.T) {
	service, _ := initTestForUserService(t)

	request := buildRandomCreateUserRequest()
	request.FirstName = ""
	resp, err := service.SaveUser(request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "first name cannot be null or empty")
}

func TestLastNameValidationErrorForRegistration(t *testing.T) {
	service, _ := initTestForUserService(t)

	request := buildRandomCreateUserRequest()
	request.LastName = ""
	resp, err := service.SaveUser(request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "last name cannot be null or empty")
}

func TestEmailValidationErrorForRegistration(t *testing.T) {
	service, _ := initTestForUserService(t)

	request := buildRandomCreateUserRequest()
	request.Email = ""
	resp, err := service.SaveUser(request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "email cannot be null or empty")
}

func TestPasswordValidationErrorForRegistration(t *testing.T) {
	service, _ := initTestForUserService(t)

	request := buildRandomCreateUserRequest()
	request.Password = ""
	resp, err := service.SaveUser(request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "password cannot be null or empty")
}

func TestPasswordFulfillsLengthUsecaseForRegistration(t *testing.T) {
	service, _ := initTestForUserService(t)

	request := buildRandomCreateUserRequest()
	request.Password = utils.RandomString(2)
	resp, err := service.SaveUser(request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "password should be at least 5 characters long")
}

func TestValidEmailForRegistration(t *testing.T) {
	service, _ := initTestForUserService(t)

	request := buildRandomCreateUserRequest()
	request.Email = utils.RandomString(5)
	resp, err := service.SaveUser(request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "is not a valid email")
}

func TestUniqueEmailForRegistration(t *testing.T) {
	service, userRepo := initTestForUserService(t)

	request := buildRandomCreateUserRequest()
	
	userRepo.EXPECT().SaveUser(gomock.Any()).Times(1).Return(nil, &commonError.ErrUniqueConstraintViolation{Message: ""})

	resp, err := service.SaveUser(request)

	var uniqueConstraintErr *commonError.ErrUniqueConstraintViolation
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &uniqueConstraintErr)
}

func TestForUserRegistration(t *testing.T) {
	service, userRepo := initTestForUserService(t)

	request := buildRandomCreateUserRequest()
	user := randomUser(request.Password)
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email

	userRepo.EXPECT().SaveUser(gomock.Any()).Times(1).Return(user, nil)

	resp, err := service.SaveUser(request)

	require.Nil(t, err)
	require.NotNil(t, resp)
	require.Equal(t, request.Email, resp.Email)
	require.Equal(t, request.FirstName, resp.FirstName)
	require.Equal(t, request.LastName, resp.LastName)
	require.Empty(t, resp.HashPassword)
	require.NotEmpty(t, resp.Id)
	require.NotEmpty(t, resp.CreatedAt)
	require.NotEmpty(t, resp.UpdatedAt)
}

func buildRandomCreateUserRequest() *domain.CreateUserRequest {
	return &domain.CreateUserRequest{
		FirstName: utils.RandomFirstName(),
		LastName:  utils.RandomLastName(),
		Email:     utils.RandomEmail(),
		Password:  utils.RandomString(10),
	}
}
