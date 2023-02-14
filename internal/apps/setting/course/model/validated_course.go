package model

import (
	"time"
)

// ValidatedCourse は検証済みの講習情報です
type ValidatedCourse struct {
	Title             string
	Description       *string
	ThumbnailImage    *ThumbnailImage
	IsRemoveThumbnail bool
	IsRequired        bool
	CategoryId        *uint32
	From              *time.Time
	To                *time.Time
}

// ThumbnailImage はサムネイル画像
type ThumbnailImage struct {
	OriginalName string
	Name         string
	Content      []byte
}
