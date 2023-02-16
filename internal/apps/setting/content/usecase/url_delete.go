package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewUrlDelete(
	connFactory database.ConnFactory,
	query content.Query,
	repos content.UrlDeleteRepository,
) content.UrlDeleteUseCase {
	return &urlDelete{
		connFactory: connFactory,
		query:       query,
		repos:       repos,
	}
}

type urlDelete struct {
	connFactory database.ConnFactory
	query       content.Query
	repos       content.UrlDeleteRepository
}

func (uc *urlDelete) Exec(
	ctx context.Context,
	courseId vo.CourseId,
	contentId vo.ContentId,
) error {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return errs.Wrap("[urlDelete.Exec]connFactory.Createのエラー", err)
	}

	exist, err := uc.query.ExistUrl(ctx, conn, courseId, contentId)
	if err != nil {
		return errs.Wrap("[urlDelete.Exec]query.ExistCourseのエラー", err)
	} else if !exist {
		return errs.NewNotFound("URLコンテンツが存在しません")
	}

	if err = uc.repos.Delete(ctx, conn, contentId); err != nil {
		return errs.Wrap("[urlDelete.Exec]repos.Deleteのエラー", err)
	}
	return nil
}
