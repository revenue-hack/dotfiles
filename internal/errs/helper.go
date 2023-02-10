package errs

import (
	"fmt"
)

// Wrap はerrorをWrapして返します
func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// GetOriginalMessage はWrap前のエラーメッセージを返却します
func GetOriginalMessage(err error) string {
	e, ok := err.(interface {
		Unwrap() error
	})
	// Unwrapを実装していないか、Unwrapがnilを返す（最上位）の場合は自身のエラーを返す
	if !ok || e.Unwrap() == nil {
		return err.Error()
	}
	return GetOriginalMessage(e.Unwrap())
}

// ErrorsOrNil はErrorsにエラーが存在する場合はerrorを、存在しない場合はnilを返却します
func ErrorsOrNil(ers *Errors) error {
	if ers.HasError() {
		return ers
	}
	return nil
}

// ErrorsOrNilWithValue はErrorsにエラーが存在する場合はnil, errorを、存在しない場合は any, nilを返却します
func ErrorsOrNilWithValue[T any](v T, ers *Errors) (*T, error) {
	if ers.HasError() {
		return nil, ers
	}
	return &v, nil
}
