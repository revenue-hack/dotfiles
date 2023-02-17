package errs

import (
	"errors"
	"fmt"
	"strings"
)

type Errors []string

func NewErrors() *Errors {
	return &Errors{}
}

func (e *Errors) Error() string {
	return strings.Join(*e, "\n")
}

func (e *Errors) Errors() []string {
	return *e
}

func (e *Errors) HasError() bool {
	return len(*e) > 0
}

func (e *Errors) Add(format string, args ...any) {
	*e = append(*e, fmt.Sprintf(format, args...))
}

func (e *Errors) AddError(err error) {
	if err == nil {
		return
	}
	*e = append(*e, err.Error())
}

func (e *Errors) Append(ers *Errors) {
	if ers == nil {
		return
	}
	*e = append(*e, *ers...)
}

func (*Errors) Is(err error) bool {
	var e *Errors
	return errors.As(err, &e)
}
