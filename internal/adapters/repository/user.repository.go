package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) SaveUser(user *domain.User) (*domain.User, error) {
	query := `insert into user (first_name, last_name, email, hash_password, role, created_at, updated_at)
	values (?, ?, ?, ?, ?, ?, ?)
	`
	res, err := r.db.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.HashPassword,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt)

	me, ok := err.(*mysql.MySQLError)
	if !ok && me != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	if me != nil && me.Number == 1062 {
		return nil, &commonError.ErrUniqueConstraintViolation{Message: fmt.Sprintf("user with email %s already exists", user.Email)}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, &commonError.ErrInternalServer{Message: err.Error()}
	}

	user.Id = id
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	query := "select id, first_name, last_name, email, role, hash_password, created_at, updated_at from user where email = ?"
	rows, err := r.db.Query(query, email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := new(domain.User)
		err = rows.Scan(&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Role,
			&user.HashPassword,
			&user.CreatedAt,
			&user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, &commonError.ErrNotFound{Message: fmt.Sprintf("%s not found", email)}
}