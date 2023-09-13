package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

const (
	SELECT_POLL_INFO_LOC        string = "./resources/sql/SelectPollInfo.sql"
	SELECT_POLL_WITH_PAGINATION string = "./resources/sql/SelectPollWithPagination.sql"
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

func (r *pollMysqlRepository) FindAllPoll(filter domain.PollListFilter) (*domain.PollPaginationDetails, error) {
	result := &domain.PollPaginationDetails{
		Data: []domain.PollWithOrganizerInfo{},
	}
	query, err := utils.LoadResourceAsString(SELECT_POLL_WITH_PAGINATION)
	if err != nil {
		return nil, err
	}

	countQuery := "select count(p.id) from poll p left join poll_organizer po on p.id = po.fk_poll_id where %s"

	args := make([]interface{}, 0)
	var condition string = "1 = 1 "
	if filter.States != nil && len(filter.States) > 0 {
		placeholder := ""
		for _, data := range filter.States {
			placeholder += "?,"
			args = append(args, data)
		}
		placeholder = strings.TrimSuffix(placeholder, ",")
		condition += fmt.Sprintf("and p.state in (%s) ", placeholder)
	}

	if filter.OrganizerId > 0 {
		condition += "and po.fk_organizer_id = ? "
		args = append(args, filter.OrganizerId)
	}

	if filter.FilterPrimaryOwner != nil {
		condition += "and po.primary_organizer = ? "
		args = append(args, filter.FilterPrimaryOwner)
	}

	offset := (filter.Page - 1) * filter.Limit
	finalSelectQuery := fmt.Sprintf(query, condition, "p."+filter.SortBy, filter.Sort, filter.Limit, offset)
	log.Println("final select query: ", finalSelectQuery)
	finalCountRowsQuery := fmt.Sprintf(countQuery, condition)
	log.Println("final count query: ", finalCountRowsQuery)

	selectStmt, err := r.db.Prepare(finalSelectQuery)
	if err != nil {
		log.Println("invalid query", err)
		return nil, err
	}
	defer selectStmt.Close()

	countRowsStmt, err := r.db.Prepare(finalCountRowsQuery)
	if err != nil {
		log.Println("invalid query", err)
		return nil, err
	}
	defer countRowsStmt.Close()

	var count int
	countResult := countRowsStmt.QueryRow(args...)
	if err := countResult.Scan(&count); err != nil {
		log.Println("No any user record in database", err)
		return result, nil
	}

	rows, err := selectStmt.Query(args...)
	if err != nil {
		log.Println("No data present")
		return result, nil
	}
	defer rows.Close()

	result.Page = int(filter.Page)
	result.Total = count
	for rows.Next() {
		poll := domain.PollWithOrganizerInfo{}
		rows.Scan(
			&poll.Id,
			&poll.Title,
			&poll.Description,
			&poll.StartsAt,
			&poll.EndsAt,
			&poll.State,
			&poll.CreatedAt,
			&poll.UpdatedAt,
			&poll.FullName,
			&poll.Email,
		)
		result.Data = append(result.Data, poll)
	}
	result.Size = len(result.Data)
	return result, nil
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
