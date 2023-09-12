package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type pollMysqlRepository struct {
	db *sql.DB
}

func NewPollMysqlRepository(db *sql.DB) ports.PollRepository {
	return &pollMysqlRepository{db: db}
}

func (r *pollMysqlRepository) SavePoll(poll *domain.Poll) (*domain.Poll, error) {
	query := "insert into poll (title, description, starts_at, ends_at, created_at, updated_at) values (?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, poll.Title, poll.Description, poll.StartsAt, poll.EndsAt, poll.CreatedAt, poll.UpdatedAt)

	me, ok := err.(*mysql.MySQLError)
	if !ok && me != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	if me != nil && me.Number == 1062 {
		return nil, &commonError.ErrUniqueConstraintViolation{Message: fmt.Sprintf("poll with title %s already exists", poll.Title)}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	poll.Id = id
	return poll, nil
}

func (r *pollMysqlRepository) FindPollById(id int64) (*domain.Poll, error) {
	return nil, nil
}

func (r *pollMysqlRepository) Init() error {
	query := `
		create table if not exists poll (
			id int not null auto_increment primary key,
			title varchar(100) not null unique,
			description longtext,
			state varchar(50) not null default "PREPARED",
			starts_at datetime not null,
			ends_at datetime not null,
			created_at timestamp not null,
			updated_at timestamp not null
		)
	`
	_, err := r.db.Exec(query)
	if err != nil {
		return &commonError.ErrInternalServer{Message: err.Error()}
	}

	return nil
}
