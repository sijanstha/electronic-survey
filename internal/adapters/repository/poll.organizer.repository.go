package repository

import (
	"database/sql"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type pollOrganizerRepository struct {
	db *sql.DB
}

func NewPollOrganizerRepository(db *sql.DB) ports.PollOrganizerRepository {
	return &pollOrganizerRepository{db}
}

func (r *pollOrganizerRepository) SavePollOrganizer(req *domain.PollOrganizer) (*domain.PollOrganizer, error) {
	query := `insert into poll_organizer(fk_poll_id, fk_organizer_id, primary_organizer, created_at, updated_at)
	values (?, ?, ?, ?, ?)
	`
	res, err := r.db.Exec(query,
		req.PollId,
		req.OrganizerId,
		req.PrimaryOrganizer,
		req.CreatedAt,
		req.UpdatedAt)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	req.Id = id
	return req, nil
}