package services

import (
	"context"
	"fmt"
	"log"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
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

func (s *participantListService) UpdateParticipantList(ctx context.Context, req *domain.UpdateParticipantList) (*domain.ParticipantList, error) {
	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == 0 {
		return nil, commonError.NewErrBadRequest("invalid participant list id")
	}

	pl, err := s.GetParticipantListById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	uniqueEmails := []string{}
	if len(req.Emails) > 0 {
		uniqueEmails = append(uniqueEmails, pl.Emails...)
		uniqueEmailMap := make(map[string]bool)
		for  _, ue := range uniqueEmails {
			uniqueEmailMap[ue] = true
		}

		for _, email := range req.Emails {
			if !utils.IsValidEmail(email) {
				return nil, commonError.NewErrBadRequest(fmt.Sprintf("invalid email provided: %s", email))
			}

			if _, ok := uniqueEmailMap[email]; !ok {
				uniqueEmailMap[email] = true
				uniqueEmails = append(uniqueEmails, email)
			}
		}
	}

	if req.Name == pl.Name {
		req.Name = ""
	}
	toUpdateParticipantList := domain.NewParticipantList(req.Id, req.Name, uniqueEmails, true)
	toUpdateParticipantList.OrganizerId = claims.Id
	return s.repo.UpdateParticipantList(toUpdateParticipantList)
}

func (s *participantListService) GetParticipantListById(ctx context.Context, id int64) (*domain.ParticipantList, error) {
	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	filter := &domain.ParticipantListFilter{
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

	filter := &domain.ParticipantListFilter{
		Name:        name,
		OrganizerId: claims.Id,
	}
	return s.repo.FindParticipantList(filter)
}

func (s *participantListService) GetParticipantList(ctx context.Context, req *domain.ParticipantListFilter) (*domain.ParticipantPaginationDetails, error) {
	if err := req.Validate(); err != nil {
		return nil, commonError.NewErrBadRequest(err.Error())
	}

	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	req.OrganizerId = claims.Id
	return s.repo.FindAllParticipantList(req)
}
