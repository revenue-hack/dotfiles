package elearning

import (
	"time"

	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
)

// GetFrom は講習の開催期間（開始）の値を返します
func GetFrom(course *entity.Course) *time.Time {
	if !course.HasSchedule() {
		return nil
	}
	// CourseSchedulesが存在する場合はELearningScheduleは必ず存在する
	return course.CourseSchedules[0].ELearningSchedule.From
}

// GetTo は講習の開催期間（終了）の値を返します
func GetTo(course *entity.Course) *time.Time {
	if !course.HasSchedule() {
		return nil
	}
	// CourseSchedulesが存在する場合はELearningScheduleは必ず存在する
	return course.CourseSchedules[0].ELearningSchedule.To
}
