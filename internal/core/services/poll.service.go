package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
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
	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
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

func (s *pollService) UpdatePoll(ctx context.Context, req *domain.UpdatePollRequest) (*domain.Poll, error) {
	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("claims: ", claims.ID)
	if req.Id == 0 {
		return nil, commonError.NewErrBadRequest("invalid poll id")
	}

	poll, err := s.GetPollById(req.Id)
	if err != nil {
		return nil, err
	}

	for _, pollOrganizer := range poll.PollOrganizers {
		if pollOrganizer.Id == claims.Id {
			break
		}
		return nil, &commonError.ErrUnauthorized{Message: fmt.Sprintf("you are not authorized to update this poll: %d, please contact poll owner", req.Id)}
	}

	if poll.State != domain.PREPARED {
		return nil, commonError.NewErrBadRequest(fmt.Sprintf("poll with state %s is not allowed for modification", poll.State))
	}

	// input validations
	var startsAt, endsAt time.Time
	if req.StartsAt != "" && len(req.StartsAt) > 0 {
		startsAt, err = utils.ParseDateTime(req.StartsAt)
		if err != nil {
			return nil, commonError.NewErrBadRequest(err.Error())
		}

		if req.EndsAt == "" {
			return nil, commonError.NewErrBadRequest("poll end date not provided")
		}

		endsAt, err = utils.ParseDateTime(req.EndsAt)
		if err != nil {
			return nil, commonError.NewErrBadRequest(err.Error())
		}
	}

	if req.EndsAt != "" && len(req.EndsAt) > 0 {
		endsAt, err = utils.ParseDateTime(req.EndsAt)
		if err != nil {
			return nil, commonError.NewErrBadRequest(err.Error())
		}

		if req.StartsAt == "" {
			return nil, commonError.NewErrBadRequest("poll start date not provided")
		}

		startsAt, err = utils.ParseDateTime(req.StartsAt)
		if err != nil {
			return nil, commonError.NewErrBadRequest(err.Error())
		}
	}

	if !startsAt.IsZero() && !endsAt.IsZero() {
		if startsAt.Before(time.Now().UTC()) {
			return nil, commonError.NewErrBadRequest("start date cannot be of past date")
		}

		if endsAt.Before(time.Now().UTC()) {
			return nil, commonError.NewErrBadRequest("end date cannot be of past date")
		}

		if endsAt.Equal(startsAt) || endsAt.Before(startsAt) {
			return nil, commonError.NewErrBadRequest("end date should be after start date")
		}
	}

	toUpdatePoll := domain.NewPoll(req.Id, req.Title, req.Description, startsAt, endsAt, req.Timezone, "", true)

	return s.pollRepo.UpdatePoll(toUpdatePoll)
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

func (s *pollService) GetAllPoll(ctx context.Context, req domain.PollListFilter) (*domain.PollPaginationDetails, error) {
	if err := req.Validate(); err != nil {
		return nil, commonError.NewErrBadRequest(err.Error())
	}

	claims, err := getCurrentLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	req.OrganizerId = claims.Id
	return s.pollRepo.FindAllPoll(req)
}
