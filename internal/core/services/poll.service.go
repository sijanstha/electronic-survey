package services

import (
	"context"
	"log"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type pollService struct {
	pollRepo          ports.PollRepository
	pollOrganizerRepo ports.PollOrganizerRepository
}

func NewPollService(repo ports.PollRepository, pollOrganizerRepo ports.PollOrganizerRepository) ports.PollService {
	return &pollService{
		pollRepo:          repo,
		pollOrganizerRepo: pollOrganizerRepo,
	}
}

func (s *pollService) SavePoll(ctx context.Context, req *domain.CreatePollRequest) (*domain.Poll, error) {
	principal := ctx.Value("principal")
	if principal == nil {
		return nil, &commonError.ErrUnauthorized{Message: "unauthorized"}
	}

	claims, ok := principal.(*domain.Claims)
	if !ok {
		return nil, &commonError.ErrInternalServer{Message: "couldn't parse jwt claims"}
	}

	log.Println("claims: ", claims.ID)

	poll, err := req.Validate()
	if err != nil {
		return nil, err
	}

	savedPoll, err := s.pollRepo.SavePoll(poll)
	if err != nil {
		return nil, err
	}

	// how to use transaction in service layer???
	pollOrganizer := domain.NewPollOrganizer(claims.Id, savedPoll.Id, true)
	_, err = s.pollOrganizerRepo.SavePollOrganizer(pollOrganizer)
	if err != nil {
		return nil, err
	}

	return savedPoll, nil
}

func (s *pollService) GetPollById(id int64) (*domain.PollInfo, error) {
	filter := domain.PollFilter{
		Id: id,
	}

	poll, err := s.pollRepo.FindPoll(filter)
	if err != nil {
		return nil, err
	}

	return poll, nil
}

func (s *pollService) GetAllPoll(req domain.PollListFilter) (*domain.PollPaginationDetails, error) {
	if err := req.Validate(); err != nil {
		return nil, commonError.NewErrBadRequest(err.Error())
	}

	return s.pollRepo.FindAllPoll(req)
}
