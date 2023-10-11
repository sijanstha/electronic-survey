package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"

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

	existingParticipantList, _ := r.FindParticipantList(&domain.ParticipantListFilter{
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

func (r *participantListRepository) UpdateParticipantList(req *domain.ParticipantList) (*domain.ParticipantList, error) {
	query := "update participant_list set %s where id = ?"

	args := make([]interface{}, 0)
	var values string
	if req.Name != "" && len(req.Name) > 0 {
		existingParticipantList, _ := r.FindParticipantList(&domain.ParticipantListFilter{
			Name:        req.Name,
			OrganizerId: req.OrganizerId,
		})

		if existingParticipantList != nil {
			return nil, &commonError.ErrUniqueConstraintViolation{Message: fmt.Sprintf("participant list with name %s already exists", req.Name)}
		}
		values += "name=?,"
		args = append(args, req.Name)
	}

	if len(req.Emails) > 0 {
		values += "emails=?,"
		jsonEmails, err := json.Marshal(req.Emails)
		if err != nil {
			return nil, &commonError.ErrInternalServer{Message: err.Error()}
		}
		args = append(args, jsonEmails)
	}

	if !req.UpdatedAt.IsZero() {
		values += "updated_at=?,"
		args = append(args, req.UpdatedAt)
	}

	values = strings.TrimSuffix(values, ",")
	args = append(args, req.Id)

	finalUpdateQuery := fmt.Sprintf(query, values)
	log.Println("Participant list Update Query: ", finalUpdateQuery)

	_, err := r.db.Exec(finalUpdateQuery, args...)
	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	return req, nil
}

func (r *participantListRepository) FindAllParticipantList(req *domain.ParticipantListFilter) (*domain.ParticipantPaginationDetails, error) {
	result := domain.ParticipantPaginationDetails{}
	query := "select id, name, emails, fk_organizer_id, created_at, updated_at from participant_list where %v order by %v %v limit %v offset %v"
	countQuery := "select count(id) from participant_list where %v"

	args := make([]interface{}, 0)
	var condition string = "1 = 1 "
	if req.OrganizerId > 0 {
		condition += "and fk_organizer_id = ? "
		args = append(args, req.OrganizerId)
	}

	if req.SearchParameter != "" && len(req.SearchParameter) > 0 {
		condition += fmt.Sprintf("and replace(concat(name, emails), ' ', '') like replace('%%%s%%', ' ', '') ", req.SearchParameter)
	}

	offset := (req.Page - 1) * req.Limit
	finalSelectQuery := fmt.Sprintf(query, condition, req.SortBy, req.Sort, req.Limit, offset)
	log.Println("final participant list select query: ", finalSelectQuery)
	finalCountRowsQuery := fmt.Sprintf(countQuery, condition)
	log.Println("final participant list count query: ", finalCountRowsQuery)

	var count int
	countResult := r.db.QueryRow(finalCountRowsQuery, args...)
	if err := countResult.Scan(&count); err != nil {
		log.Println("No any participant list record in database", err)
		return nil, &commonError.ErrNotFound{Message: "participant list not found"}
	}

	rows, err := r.db.Query(finalSelectQuery, args...)
	if err != nil {
		log.Println("No any participant list record in database", err)
		return nil, &commonError.ErrNotFound{Message: "participant list not found"}
	}
	defer rows.Close()

	result.Page = int(req.Page)
	result.Total = count
	for rows.Next() {
		pl := domain.ParticipantList{}
		var rawEmailInBytes []byte
		rows.Scan(&pl.Id, &pl.Name, &rawEmailInBytes, &pl.OrganizerId, &pl.CreatedAt, &pl.UpdatedAt)
		var emails []string
		err := json.Unmarshal(rawEmailInBytes, &emails)

		if err != nil {
			return nil, &commonError.ErrNotFound{Message: err.Error()}
		}

		pl.Emails = emails
		result.Data = append(result.Data, pl)
	}
	result.Size = int(req.Limit)
	totalPages := int(math.Ceil(float64(result.Total) / float64(result.Size)))
	result.TotalPages = totalPages
	return &result, nil
}

func (r *participantListRepository) FindParticipantList(filter *domain.ParticipantListFilter) (*domain.ParticipantList, error) {
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
