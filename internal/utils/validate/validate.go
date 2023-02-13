package validate

import (
	"unicode/utf8"

	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

// StringRequired 必須入力の文字列を検証します
func StringRequired(fieldName string, val *string, maxLength int) error {
	if val == nil || *val == "" {
		return errs.NewInvalidParameter("%sは必須です", fieldName)
	}
	if utf8.RuneCount([]byte(*val)) > maxLength {
		return errs.NewInvalidParameter("%sは%d文字以内で入力してください", fieldName, maxLength)
	}
	return nil
}

// StringOptional 必須入力の文字列を検証します
func StringOptional(fieldName string, val *string, maxLength int) error {
	if val == nil || *val == "" {
		return nil
	}
	if utf8.RuneCount([]byte(*val)) > maxLength {
		return errs.NewInvalidParameter("%sは%d文字以内で入力してください", fieldName, maxLength)
	}
	return nil
}
