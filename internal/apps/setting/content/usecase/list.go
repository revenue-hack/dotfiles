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
	query content.Query,
	listQuery content.ListQuery,
) content.ListUseCase {
	return &list{
		connFactory: connFactory,
		query:       query,
		listQuery:   listQuery,
	}
}

type list struct {
	connFactory database.ConnFactory
	query       content.Query
	listQuery   content.ListQuery
}

func (uc *list) Exec(ctx context.Context, courseId vo.CourseId) (*content.ListOutput, error) {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[urlCreate.Exec]connFactory.Createのエラー", err)
	}

	exist, err := uc.query.ExistCourse(ctx, conn, courseId)
	if err != nil {
		return nil, errs.Wrap("[urlCreate.Exec]query.ExistCourseのエラー", err)
	} else if !exist {
		return nil, errs.NewNotFound("講習が存在しません")
	}

	contents, err := uc.listQuery.GetAll(ctx, conn, courseId)
	if err != nil {
		return nil, errs.Wrap("[urlCreate.Exec]listQuery.GetAllのエラー", err)
	}

	return &content.ListOutput{
		Contents: contents,
	}, nil
}
