package commonError

type ErrUniqueConstraintViolation struct {
	Message string
}

func (e *ErrUniqueConstraintViolation) Error() string {
	return e.Message
}

type ErrInternalServer struct {
	Message string
}

func (e *ErrInternalServer) Error() string {
	return e.Message
}

type ErrBadRequest struct {
	Message string
}

func (e *ErrBadRequest) Error() string {
	return e.Message
}

type ErrUnauthorized struct {
	Message string
}

func (e *ErrUnauthorized) Error() string {
	return e.Message
}

type ErrNotFound struct {
	Message string
}

func (e *ErrNotFound) Error() string {
	return e.Message
}