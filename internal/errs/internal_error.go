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

func (e *InternalError) Error() string {
	return e.msg
}

func (*InternalError) Is(err error) bool {
	var e *InternalError
	return errors.As(err, &e)
}
