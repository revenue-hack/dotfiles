package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewQuery() content.Query {
	return &query{}
}

type query struct{}

func (r *query) ExistCourse(
	ctx context.Context,
	conn *database.Conn,
	courseId vo.CourseId,
) (bool, error) {
	query := conn.DB().
		Where("id = ?", courseId.Value())

	exist, err := database.Exist[entity.Course](ctx, query)
	if err != nil {
		return exist, errs.Wrap("[query.ExistCourse]database.Existのエラー", err)
	}
	return exist, nil
}

func (r *query) ExistUrl(
	ctx context.Context,
	conn *database.Conn,
	courseId vo.CourseId,
	contentId vo.ContentId,
) (bool, error) {
	query := conn.DB().
		Where("id = ?", contentId.Value()).
		Where("course_id = ?", courseId.Value()).
		Where("content_type = ? ", entity.ContentTypeUrl)

	exist, err := database.Exist[entity.Content](ctx, query)
	if err != nil {
		return exist, errs.Wrap("[query.ExistUrl]database.Existのエラー", err)
	}
	return exist, nil
}
