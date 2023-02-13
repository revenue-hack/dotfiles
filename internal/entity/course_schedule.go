package entity

import (
	"time"
)

type CourseSchedule struct {
	Id        uint32
	CourseId  uint32
	CreatedAt time.Time
	CreatedBy uint32
	UpdatedAt time.Time
	UpdatedBy uint32

	// relations

	ELearningSchedule ELearningSchedule
}

type CourseSchedules = []CourseSchedule
