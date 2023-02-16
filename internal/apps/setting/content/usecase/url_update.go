package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewUrlUpdate(
	connFactory database.ConnFactory,
	query content.Query,
	repos content.UrlUpdateRepository,
	service content.UrlService,
) content.UrlUpdateUseCase {
	return &urlUpdate{
		connFactory: connFactory,
		query:       query,
		repos:       repos,
		service:     service,
	}
}

type urlUpdate struct {
	connFactory database.ConnFactory
	query       content.Query
	repos       content.UrlUpdateRepository
	service     content.UrlService
}

func (uc *urlUpdate) Exec(
	ctx context.Context,
	courseId vo.CourseId,
	contentId vo.ContentId,
	in content.UrlInput,
) error {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return errs.Wrap("[urlUpdate.Exec]connFactory.Createのエラー", err)
	}

	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return errs.Wrap("[urlUpdate.Exec]ctxt.AuthenticatedUserのエラー", err)
	}

	exist, err := uc.query.ExistUrl(ctx, conn, courseId, contentId)
	if err != nil {
		return errs.Wrap("[urlUpdate.Exec]query.ExistCourseのエラー", err)
	} else if !exist {
		return errs.NewNotFound("URLコンテンツが存在しません")
	}

	validUrl, err := uc.service.NewValidatedUrl(in)
	if err != nil {
		return errs.Wrap("[urlUpdate.Exec]service.NewValidatedUrlのエラー", err)
	}

	if err = uc.repos.Update(ctx, conn, authedUser, contentId, validUrl); err != nil {
		return errs.Wrap("[urlUpdate.Exec]repos.Updateのエラー", err)
	}
	return nil
}
