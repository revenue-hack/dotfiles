//go:build !feature_test

package timer

import (
	"time"
)

func Now() time.Time {
	return time.Now().In(loc)
}

func Parse(format, value string) (time.Time, error) {
	return time.ParseInLocation(format, value, loc)
}
