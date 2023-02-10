//go:build feature_test

package timer

import (
	"time"
)

func Now() time.Time {
	return time.Date(2023, 4, 3, 12, 34, 56, 0, loc)
}

func Parse(format, value string) (time.Time, error) {
	return time.ParseInLocation(format, value, loc)
}
