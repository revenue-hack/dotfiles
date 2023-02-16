package model

import (
	"time"
)

type Course struct {
	Id           uint32
	Title        string
	Thumbnail    *Thumbnail
	CategoryName *string
	ExpireAt     *time.Time
	Recommend    *uint32
	IsRequired   bool
	IsFixed      bool
}

type Thumbnail struct {
	Name string
	Url  string
}
