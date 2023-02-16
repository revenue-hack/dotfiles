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
		return exist, errs.Wrap("[urlCreateRepository.ExistCourse]database.Existのエラー", err)
	}
	return exist, nil
}
