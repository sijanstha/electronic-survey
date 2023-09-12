package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

const (
	SELECT_POLL_INFO_LOC string = "./resources/sql/SelectPollInfo.sql"
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

func (r *pollMysqlRepository) FindPoll(filter domain.PollFilter) (*domain.PollInfo, error) {
	query, err := utils.LoadResourceAsString(SELECT_POLL_INFO_LOC)
	if err != nil {
		return nil, err
	}

	condition := "1=1"
	if filter.Id > 0 {
		condition += " and p.id = " + fmt.Sprintf("%d", filter.Id)
	}

	if filter.Title != "" && len(filter.Title) > 0 {
		condition += " and p.title = '" + filter.Title + "'"
	}

	finalSelectQuery := fmt.Sprintf(query, condition)
	log.Println("query: ", finalSelectQuery)

	pollInfoJson := new(string)
	result := r.db.QueryRow(finalSelectQuery)

	if err := result.Scan(pollInfoJson); err != nil {
		log.Println(err)
		return nil, err
	}

	pollDatabaseView := new(domain.PollDatabaseView)
	err = json.Unmarshal([]byte(*pollInfoJson), pollDatabaseView)
	if err != nil {
		return nil, err
	}

	if pollDatabaseView.Id == 0 {
		return nil, &commonError.ErrNotFound{Message: "poll not found"}
	}

	return pollDatabaseView.ToPollInfo(), nil
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
