package entity

import (
	"time"
)

type Url struct {
	Id        uint32
	ContentId uint32
	Title     string
	Url       string
	CreatedAt time.Time
	CreatedBy uint32
	UpdatedAt time.Time
	UpdatedBy uint32
}

type Urls = []Url
