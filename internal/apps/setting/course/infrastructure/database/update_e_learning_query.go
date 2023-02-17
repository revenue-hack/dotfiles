package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewUpdateELearningQuery() course.UpdateQuery {
	return &updateELearningQuery{}
}

type updateELearningQuery struct{}

func (*updateELearningQuery) ExistCourse(
	ctx context.Context,
	conn *database.Conn,
	courseId vo.CourseId,
) (bool, error) {
	query := conn.DB().
		Where("id = ?", courseId.Value()).
		Where("course_type = ?", entity.CourseTypeELearning)

	exist, err := database.Exist[entity.Course](ctx, query)
	if err != nil {
		return exist, errs.Wrap("[updateELearningQuery.ExistCourse]database.Existのエラー", err)
	}
	return exist, nil
}

func (*updateELearningQuery) ExistCategory(
	ctx context.Context,
	conn *database.Conn,
	categoryId uint32,
) (bool, error) {
	exist, err := database.ExistById[entity.Category](ctx, conn.DB(), categoryId)
	if err != nil {
		return exist, errs.Wrap("[updateELearningQuery.ExistCategory]database.ExistByIdのエラー", err)
	}
	return exist, nil
}
