package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type participantListRepository struct {
	db *sql.DB
}

func NewParticipantListRepository(db *sql.DB) ports.ParticipantListRepository {
	return &participantListRepository{db: db}
}

func (r *participantListRepository) SaveParticipantList(req *domain.ParticipantList) (*domain.ParticipantList, error) {
	jsonEmails, err := json.Marshal(req.Emails)
	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	existingParticipantList, _ := r.FindParticipantList(&domain.ParticipantFilter{
		Name:        req.Name,
		OrganizerId: req.OrganizerId,
	})

	if existingParticipantList != nil {
		return nil, &commonError.ErrUniqueConstraintViolation{Message: fmt.Sprintf("participant list with name %s already exists", req.Name)}
	}

	query := "insert into participant_list(name, emails, fk_organizer_id, created_at, updated_at) values (?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, req.Name, string(jsonEmails), req.OrganizerId, req.CreatedAt, req.UpdatedAt)

	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	req.Id = id
	return req, nil
}

func (r *participantListRepository) FindAllParticipantList(req *domain.ParticipantListFilter) (domain.ParticipantPaginationDetails, error) {
	return domain.ParticipantPaginationDetails{}, nil
}

func (r *participantListRepository) FindParticipantList(filter *domain.ParticipantFilter) (*domain.ParticipantList, error) {
	query := "select id, name, emails, fk_organizer_id, created_at, updated_at from participant_list where %s"

	condition := "1=1"
	args := make([]interface{}, 0)
	if filter.Id > 0 {
		condition += " and id = ?"
		args = append(args, filter.Id)
	}

	if filter.Name != "" && len(filter.Name) > 0 {
		condition += " and name = ?"
		args = append(args, filter.Name)
	}

	if filter.OrganizerId > 0 {
		condition += " and fk_organizer_id = ?"
		args = append(args, filter.OrganizerId)
	}

	finalSelectQuery := fmt.Sprintf(query, condition)
	log.Println("query: ", finalSelectQuery)

	pl := domain.ParticipantList{}
	result := r.db.QueryRow(finalSelectQuery, args...)

	var rawEmailInBytes []byte
	if err := result.Scan(&pl.Id, &pl.Name, &rawEmailInBytes, &pl.OrganizerId, &pl.CreatedAt, &pl.UpdatedAt); err != nil {
		return nil, &commonError.ErrNotFound{Message: "participant list not found"}
	}

	var emails []string
	err := json.Unmarshal(rawEmailInBytes, &emails)

	if err != nil {
		return nil, &commonError.ErrNotFound{Message: err.Error()}
	}

	pl.Emails = emails
	return &pl, nil
}
