package errs

import (
	"errors"
	"fmt"
)

type InternalError struct {
	msg string
}

func NewInternalError(format string, args ...any) *InternalError {
	return &InternalError{msg: fmt.Sprintf(format, args...)}
}

func (r *InternalError) Error() string {
	return r.msg
}

func (r *InternalError) Is(err error) bool {
	var e *InternalError
	return errors.As(err, &e)
}
