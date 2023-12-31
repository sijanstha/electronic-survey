package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

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
	query := "insert into poll (title, description, starts_at, ends_at, timezone, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, poll.Title, poll.Description, poll.StartsAt, poll.EndsAt, poll.Timezone, poll.CreatedAt, poll.UpdatedAt)

	me, ok := err.(*mysql.MySQLError)
	if !ok && me != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	if me != nil && me.Number == 1062 {
		return nil, &commonError.ErrUniqueConstraintViolation{Message: fmt.Sprintf("poll with title %s already exists", poll.Title)}
	}

	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	poll.Id = id
	return poll, nil
}

func (r *pollMysqlRepository) UpdatePoll(poll *domain.Poll) (*domain.Poll, error) {
	query := "update poll set %s where id = ?"

	args := make([]interface{}, 0)
	var values string
	if poll.Title != "" && len(poll.Title) > 0 {
		values += "title=?,"
		args = append(args, poll.Title)
	}

	if poll.Description != "" && len(poll.Description) > 0 {
		values += "description=?,"
		args = append(args, poll.Description)
	}

	if poll.State != "" && len(poll.State) > 0 {
		values += "state=?,"
		args = append(args, poll.State)
	}

	if poll.Timezone != "" && len(poll.Timezone) > 0 {
		values += "timezone=?,"
		args = append(args, poll.Timezone)
	}

	if !poll.StartsAt.IsZero() {
		values += "starts_at=?,"
		args = append(args, poll.StartsAt)
	}

	if !poll.EndsAt.IsZero() {
		values += "ends_at=?,"
		args = append(args, poll.EndsAt)
	}

	if !poll.UpdatedAt.IsZero() {
		values += "updated_at=?,"
		args = append(args, poll.UpdatedAt)
	}

	values = strings.TrimSuffix(values, ",")
	args = append(args, poll.Id)

	finalUpdateQuery := fmt.Sprintf(query, values)
	log.Printf("Poll Update Query: %s", finalUpdateQuery)

	stmt, err := r.db.Prepare(finalUpdateQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	me, ok := err.(*mysql.MySQLError)
	if !ok && me != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	if me != nil && me.Number == 1062 {
		return nil, &commonError.ErrUniqueConstraintViolation{Message: fmt.Sprintf("poll with title %s already exists", poll.Title)}
	}

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
			&poll.Timezone,
			&poll.CreatedAt,
			&poll.UpdatedAt,
			&poll.FullName,
			&poll.Email,
		)
		result.Data = append(result.Data, poll)
	}
	result.Size = int(filter.Limit)
	totalPages := int(math.Ceil(float64(result.Total) / float64(result.Size)))
	result.TotalPages = totalPages
	return result, nil
}

func (r *pollMysqlRepository) FindAllPollInStartedStateInDateRange(from time.Time, to time.Time) ([]*domain.Poll, error) {
	query := `select 
			id, title, state, description, starts_at, ends_at, timezone, created_at, updated_at 
			from Poll p 
			where p.state = ? 
			and p.starts_at >= ? 
			and p.starts_at < ?`

	rows, err := r.db.Query(query, domain.STARTED, from, to)
	if err != nil {
		return nil, err
	}

	polls := make([]*domain.Poll, 0)
	for rows.Next() {
		poll, err := r.scanPollRow(rows)
		if err != nil {
			return nil, err
		}
		polls = append(polls, poll)
	}

	return polls, nil
}

func (r *pollMysqlRepository) FindAllPollInVotingStateInDateRange(from time.Time, to time.Time) ([]*domain.Poll, error) {
	query := `select 
			id, title, state, description, starts_at, ends_at, timezone, created_at, updated_at 
			from Poll p 
			where p.state = ? 
			and p.ends_at >= ? 
			and p.ends_at < ?`

	rows, err := r.db.Query(query, domain.VOTING, from, to)
	if err != nil {
		return nil, err
	}

	polls := make([]*domain.Poll, 0)
	for rows.Next() {
		poll, err := r.scanPollRow(rows)
		if err != nil {
			return nil, err
		}
		polls = append(polls, poll)
	}

	return polls, nil
}

func (r *pollMysqlRepository) scanPollRow(rows *sql.Rows) (*domain.Poll, error) {
	poll := new(domain.Poll)
	err := rows.Scan(&poll.Id,
		&poll.Title,
		&poll.State,
		&poll.Description,
		&poll.StartsAt,
		&poll.EndsAt,
		&poll.Timezone,
		&poll.CreatedAt,
		&poll.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return poll, nil
}
