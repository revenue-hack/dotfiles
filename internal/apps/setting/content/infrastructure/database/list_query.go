package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewListQuery() content.ListQuery {
	return &listQuery{}
}

type listQuery struct{}

func (r *listQuery) ExistCourse(
	ctx context.Context,
	conn *database.Conn,
	courseId vo.CourseId,
) (bool, error) {
	query := conn.DB().
		Where("id = ?", courseId.Value())

	exist, err := database.Exist[entity.Course](ctx, query)
	if err != nil {
		return exist, errs.Wrap("[listQuery.ExistCourse]database.Existのエラー", err)
	}
	return exist, nil
}

func (r *listQuery) GetAll(
	ctx context.Context,
	conn *database.Conn,
	courseId vo.CourseId,
) (entity.Contents, error) {
	query := conn.DB().
		Select([]string{
			"c.id",
			"c.content_type",
		}).
		Table("contents c").
		Preload("Url").
		Where("c.course_id = ?", courseId.Value()).
		Order("c.display_order ASC")

	record, err := database.GetAll[entity.Content](ctx, query)
	if err != nil {
		return nil, errs.Wrap("[listQuery.GetAll]database.GetAllのエラー", err)
	}
	return record, nil
}
