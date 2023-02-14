package model

import (
	"time"
)

// ValidatedCourse は検証済みの講習情報です
type ValidatedCourse struct {
	Title             string
	Description       *string
	Thumbnail         *Thumbnail
	IsRemoveThumbnail bool
	IsRequired        bool
	CategoryId        *uint32
	From              *time.Time
	To                *time.Time
}

type Thumbnail struct {
	OriginalName string
	Name         string
	Content      []byte
}
