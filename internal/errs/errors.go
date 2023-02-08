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

func (r *Errors) Error() string {
	return strings.Join(*r, "\n")
}

func (r *Errors) Errors() []string {
	return *r
}

func (r *Errors) HasError() bool {
	return len(*r) > 0
}

func (r *Errors) Add(format string, args ...any) {
	*r = append(*r, fmt.Sprintf(format, args...))
}

func (r *Errors) AddError(err error) {
	if err == nil {
		return
	}
	*r = append(*r, err.Error())
}

func (r *Errors) Append(ers *Errors) {
	if ers == nil {
		return
	}
	*r = append(*r, *ers...)
}

func (r *Errors) Is(err error) bool {
	var e *Errors
	return errors.As(err, &e)
}
