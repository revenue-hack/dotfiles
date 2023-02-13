package entity

import (
	"time"
)

type Category struct {
	Id           uint32
	Name         string
	DisplayOrder uint16
	CreatedAt    time.Time
	CreatedBy    uint32
	UpdatedAt    time.Time
	UpdatedBy    uint32
}

type Categories = []Category
