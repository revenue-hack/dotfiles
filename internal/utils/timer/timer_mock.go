//go:build feature_test

package timer

import (
	"time"
)

func Now() time.Time {
	return time.Date(2022, 10, 4, 12, 34, 56, 0, loc)
}

func Parse(format, value string) (time.Time, error) {
	return time.ParseInLocation(format, value, loc)
}
