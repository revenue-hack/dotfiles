package errs

import (
	"errors"
	"fmt"
)

type NotFound struct {
	msg string
}

func NewNotFound(format string, args ...any) *NotFound {
	return &NotFound{msg: fmt.Sprintf(format, args...)}
}

func (r *NotFound) Error() string {
	return r.msg
}

func (r *NotFound) Is(err error) bool {
	var e *NotFound
	return errors.As(err, &e)
}
