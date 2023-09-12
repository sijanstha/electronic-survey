package commonError

func NewErrBadRequest(msg string) *ErrBadRequest {
	return &ErrBadRequest{Message: msg}
}