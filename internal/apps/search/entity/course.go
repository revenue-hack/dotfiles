package entity

import (
	"time"
)

type Course struct {
	Id                        uint32
	Title                     string
	ThumbnailDeliveryFileName *string
	ThumbnailOriginalFileName *string
	CategoryName              *string
	From                      *time.Time
	To                        *time.Time
	IsRequired                bool
}
