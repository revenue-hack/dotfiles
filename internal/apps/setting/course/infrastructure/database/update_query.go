package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewUpdateQuery() course.UpdateQuery {
	return &updateQuery{}
}

type updateQuery struct{}

func (r *updateQuery) ExistCourse(ctx context.Context, conn *database.Conn, courseId vo.CourseId) (bool, error) {
	exist, err := database.ExistById[entity.Course](ctx, conn.DB(), courseId.Value())
	if err != nil {
		return exist, errs.Wrap("[updateQuery.ExistCourse]database.ExistByIdのエラー", err)
	}
	return exist, nil
}

func (r *updateQuery) ExistCategory(ctx context.Context, conn *database.Conn, categoryId uint32) (bool, error) {
	exist, err := database.ExistById[entity.Category](ctx, conn.DB(), categoryId)
	if err != nil {
		return exist, errs.Wrap("[updateQuery.ExistCategory]database.ExistByIdのエラー", err)
	}
	return exist, nil
}
