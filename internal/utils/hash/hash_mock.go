//go:build feature_test

package hash

import (
	"fmt"
	"strings"
)

func Make(s string) string {
	// .を-にして固定値をつけて返すだけ
	return fmt.Sprintf("%s-hashed", strings.ReplaceAll(s, ".", "-"))
}
