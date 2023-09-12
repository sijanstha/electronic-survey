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

func (r *pollOrganizerRepository) Init() error {
	query := `
		create table if not exists poll_organizer (
			id int not null auto_increment primary key,
			fk_poll_id int not null,
			fk_organizer_id int not null,
			primary_organizer bit(1) default 0,
			created_at timestamp not null,
			updated_at timestamp not null,
			foreign key(fk_poll_id) references poll(id),
			foreign key(fk_organizer_id) references user(id)
		)
	`
	_, err := r.db.Exec(query)
	if err != nil {
		return &commonError.ErrInternalServer{Message: err.Error()}
	}

	return nil
}
