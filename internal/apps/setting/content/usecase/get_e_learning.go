package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewList(
	connFactory database.ConnFactory,
	query content.GetQuery,
) content.ListUseCase {
	return &list{
		connFactory: connFactory,
		query:       query,
	}
}

type list struct {
	connFactory database.ConnFactory
	query       content.GetQuery
}

func (uc *list) Exec(ctx context.Context, courseId vo.CourseId) (*content.GetOutput, error) {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[list.Exec]connFactory.Createのエラー", err)
	}

	contents, err := uc.query.GetAll(ctx, conn, courseId)
	if err != nil {
		return nil, errs.Wrap("[list.Exec]query.GetAllのエラー", err)
	}

	return &content.GetOutput{
		Contents: contents,
	}, nil
}
