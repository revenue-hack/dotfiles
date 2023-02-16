package entity

import (
	"time"
)

type File struct {
	Id               uint32
	ContentId        uint32
	DeliveryFileName string
	OriginalFileName string
	CreatedAt        time.Time
	CreatedBy        uint32
	UpdatedAt        time.Time
	UpdatedBy        uint32
}

type Files = []File
