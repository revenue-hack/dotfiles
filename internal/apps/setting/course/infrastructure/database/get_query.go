package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewGetQuery() course.GetQuery {
	return &getQuery{}
}

type getQuery struct{}

func (r *getQuery) GetELearning(
	ctx context.Context,
	conn *database.Conn,
	courseId vo.CourseId,
) (*entity.Course, error) {
	query := conn.DB().
		Select([]string{
			"c.id",
			"c.title",
			"c.description",
			"c.thumbnail_image_name",
			"c.is_required",
			"c.category_id",
		}).
		Table("courses c").
		Preload("CourseSchedules").
		Preload("CourseSchedules.ELearningSchedule").
		Where("c.id = ?", courseId.Value()).
		Where("c.course_type = ?", entity.CourseTypeELearning)

	record, err := database.Get[entity.Course](ctx, query)
	if err != nil {
		return nil, errs.Wrap("[getQuery.GetELearning]database.Getのエラー", err)
	}
	return record, nil
}
