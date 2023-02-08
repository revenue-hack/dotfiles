package errs

import (
	"errors"
	"fmt"
)

type InvalidParameter struct {
	msg string
}

func NewInvalidParameter(format string, args ...any) *InvalidParameter {
	return &InvalidParameter{msg: fmt.Sprintf(format, args...)}
}

func (r *InvalidParameter) Error() string {
	return r.msg
}

func (r *InvalidParameter) Is(err error) bool {
	var e *InvalidParameter
	return errors.As(err, &e)
}
