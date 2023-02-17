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

func (e *NotFound) Error() string {
	return e.msg
}

func (*NotFound) Is(err error) bool {
	var e *NotFound
	return errors.As(err, &e)
}
