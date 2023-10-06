package services

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	mockrepository "github.com/sijanstha/electronic-voting-system/internal/adapters/repository/mock"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func initTestForParticipantListService(t *testing.T) (ports.ParticipantListService, *mockrepository.MockParticipantListRepository) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockrepository.NewMockParticipantListRepository(ctrl)
	return NewParticipantListService(repo), repo
}

func TestNameValidationErrorForCreatingParticipantList(t *testing.T) {
	service, _ := initTestForParticipantListService(t)

	ctx := buildRandomContext()
	request := buildRandomCreateParticipantListRequest()
	request.Name = ""
	resp, err := service.SaveParticipantList(ctx, request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "name cannot be null or empty")
}

func TestEmailListEmptyValidationErrorForCreatingParticipantList(t *testing.T) {
	service, _ := initTestForParticipantListService(t)

	ctx := buildRandomContext()
	request := buildRandomCreateParticipantListRequest()
	request.Emails = []string{}
	resp, err := service.SaveParticipantList(ctx, request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "please provide at least one email")
}

func TestValidEmailListValidationErrorForCreatingParticipantList(t *testing.T) {
	service, _ := initTestForParticipantListService(t)

	ctx := buildRandomContext()
	request := buildRandomCreateParticipantListRequest()
	request.Emails = append(request.Emails, "invalid-email")
	resp, err := service.SaveParticipantList(ctx, request)

	var badRequestErr *commonError.ErrBadRequest
	require.Nil(t, resp)
	require.NotNil(t, err)
	require.ErrorAs(t, err, &badRequestErr)
	require.ErrorContains(t, err, "invalid email provided")
}

func buildRandomCreateParticipantListRequest() *domain.CreateParticipantListRequest {
	return &domain.CreateParticipantListRequest{
		Name:   utils.RandomString(12),
		Emails: []string{utils.RandomEmail(), utils.RandomEmail(), utils.RandomEmail()},
	}
}

func buildRandomContext() context.Context {
	return context.WithValue(context.Background(), "principal", &domain.Claims{
		Email:            utils.RandomEmail(),
		Id:               utils.RandomInt(1, 100),
		Role:             utils.RandomRole(),
		RegisteredClaims: jwt.RegisteredClaims{},
	})
}
