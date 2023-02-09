package entity

import (
	"time"
)

type Course struct {
	Id                 uint32
	Title              string
	ThumbnailImageName *string
	CategoryName       *string
	From               *time.Time
	To                 *time.Time
	IsRequired         bool
}
