package entity

import (
	"time"
)

type ELearningSchedule struct {
	Id               uint32
	CourseScheduleId uint32
	From             *time.Time
	To               *time.Time
	CreatedAt        time.Time
	CreatedBy        uint32
	UpdatedAt        time.Time
	UpdatedBy        uint32
}

type ELearningSchedules = []ELearningSchedule
