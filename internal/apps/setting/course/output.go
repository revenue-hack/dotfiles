package course

import (
	"time"
)

type GetELearningOutput struct {
	Id          uint32
	Title       string
	Description *string
	Thumbnail   *GetELearningThumbnailOutput
	IsRequired  bool
	CategoryId  *uint32
	From        *time.Time
	To          *time.Time
}

type GetELearningThumbnailOutput struct {
	Name string
	Url  string
}
