package timer

import (
	"time"
)

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Asia/Tokyo")
}
