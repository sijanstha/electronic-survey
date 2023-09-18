package domain

import (
	"fmt"
	"time"

	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

const (
	minPasswordLength int = 5
)

func (c *CreatePollRequest) Validate() (*Poll, error) {
	if c.Title == "" {
		return nil, commonError.NewErrBadRequest("title cannot be empty or null")
	}

	if c.StartsAt == "" {
		return nil, commonError.NewErrBadRequest("starts at cannot be empty or null")
	}

	if c.EndsAt == "" {
		return nil, commonError.NewErrBadRequest("ends at cannot be empty or null")
	}

	if c.Timezone == "" {
		return nil, commonError.NewErrBadRequest("timezone cannot be empty or null")
	}

	startsAt, err := utils.ParseDateTime(c.StartsAt)
	if err != nil {
		return nil, commonError.NewErrBadRequest(err.Error())
	}

	endsAt, err := utils.ParseDateTime(c.EndsAt)
	if err != nil {
		return nil, commonError.NewErrBadRequest(err.Error())
	}

	if startsAt.Before(time.Now().UTC()) {
		return nil, commonError.NewErrBadRequest("start date cannot be of past date")
	}

	if endsAt.Before(time.Now().UTC()) {
		return nil, commonError.NewErrBadRequest("end date cannot be of past date")
	}

	if endsAt.Equal(startsAt) || endsAt.Before(startsAt) {
		return nil, commonError.NewErrBadRequest("end date should be after start date")
	}

	return NewPoll(0, c.Title, c.Description, startsAt, endsAt, c.Timezone, PREPARED, false), nil
}

func (c *CreateUserRequest) Validate() (*User, error) {

	if c.FirstName == "" {
		return nil, commonError.NewErrBadRequest("first name cannot be null or empty")
	}

	if c.LastName == "" {
		return nil, commonError.NewErrBadRequest("last name cannot be null or empty")
	}

	if c.Email == "" {
		return nil, commonError.NewErrBadRequest("email cannot be null or empty")
	}

	if c.Password == "" {
		return nil, commonError.NewErrBadRequest("password cannot be null or empty")
	}

	if !utils.IsValidEmail(c.Email) {
		return nil, commonError.NewErrBadRequest(fmt.Sprintf("%s is not a valid email", c.Email))
	}

	if len(c.Password) < minPasswordLength {
		return nil, commonError.NewErrBadRequest(fmt.Sprintf("password should be at least %d characters long", minPasswordLength))
	}

	return NewUser(*c), nil
}

func (c *LoginRequest) Validate() error {
	if c.Email == "" {
		return commonError.NewErrBadRequest("email cannot be null or empty")
	}

	if c.Password == "" {
		return commonError.NewErrBadRequest("password cannot be null or empty")
	}

	if len(c.Password) < minPasswordLength {
		return commonError.NewErrBadRequest(fmt.Sprintf("password should be at least %d characters long", minPasswordLength))
	}

	return nil
}
