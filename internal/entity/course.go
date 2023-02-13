package entity

import (
	"time"
)

const (
	// CourseTypeELearning 講習区分 - e-Learning
	CourseTypeELearning uint8 = 1
	// CourseTypeGroupTraining 講習区分 - 集合研修
	CourseTypeGroupTraining uint8 = 2

	// CourseIsOptional 講習の必須フラグがOFF
	CourseIsOptional = true
	// CourseIsRequired 講習の必須フラグがON
	CourseIsRequired = true

	// CourseStatusPrivate 講習のステータス - 非公開
	CourseStatusPrivate uint8 = 1
	// CourseStatusPublic 講習のステータス - 公開
	CourseStatusPublic uint8 = 2
)

type Course struct {
	Id                 uint32
	CourseType         uint8
	Title              string
	Description        *string
	ThumbnailImageName *string
	IsRequired         bool
	CategoryId         *uint32
	Status             uint8
	CreatedAt          time.Time
	CreatedBy          uint32
	UpdatedAt          time.Time
	UpdatedBy          uint32

	// relations

	CourseSchedules CourseSchedules
}

type Courses = []Course
