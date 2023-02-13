package model

import (
	"time"
)

// ValidatedCourse は検証済みの講習情報です
type ValidatedCourse struct {
	Title          string
	Description    *string
	ThumbnailImage *string
	IsRequired     bool
	CategoryId     *uint32
	From           *time.Time
	To             *time.Time
}
