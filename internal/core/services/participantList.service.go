package services

import (
	"context"
	"log"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type participantListService struct {
	repo ports.ParticipantListRepository
}

func NewParticipantListService(repo ports.ParticipantListRepository) ports.ParticipantListService {
	return &participantListService{repo: repo}
}

func (s *participantListService) SaveParticipantList(ctx context.Context, req *domain.CreateParticipantListRequest) (*domain.ParticipantList, error) {
	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("claims: ", claims.ID)
	pl, err := req.Validate()
	if err != nil {
		return nil, err
	}

	pl.OrganizerId = claims.Id
	return s.repo.SaveParticipantList(pl)
}

func (s *participantListService) GetParticipantListById(ctx context.Context, id int64) (*domain.ParticipantList, error) {
	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	filter := &domain.ParticipantFilter{
		Id:          id,
		OrganizerId: claims.Id,
	}
	return s.repo.FindParticipantList(filter)
}

func (s *participantListService) GetParticipantListByName(ctx context.Context, name string) (*domain.ParticipantList, error) {
	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	filter := &domain.ParticipantFilter{
		Name:        name,
		OrganizerId: claims.Id,
	}
	return s.repo.FindParticipantList(filter)
}
