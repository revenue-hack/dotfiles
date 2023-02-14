package model

import (
	"time"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/domain/file"
)

// ValidatedCourse は検証済みの講習情報です
type ValidatedCourse struct {
	Title             string
	Description       *string
	Thumbnail         *file.UploadFile
	IsRemoveThumbnail bool
	IsRequired        bool
	CategoryId        *uint32
	From              *time.Time
	To                *time.Time
}
