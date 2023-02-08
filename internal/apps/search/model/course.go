package model

import (
	"time"
)

type Course struct {
	Id           uint32
	Title        string
	ThumbnailUrl string
	CategoryName *string
	ExpireAt     *time.Time
	Recommend    *uint32
	IsRequired   bool
	IsFixed      bool
}
