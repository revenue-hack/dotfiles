package course

import (
	"time"
)

type GetELearningOutput struct {
	Id           uint32
	Title        string
	Description  *string
	ThumbnailUrl *string
	IsRequired   bool
	CategoryId   *uint32
	From         *time.Time
	To           *time.Time
}
