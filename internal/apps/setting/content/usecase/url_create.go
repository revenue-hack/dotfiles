package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewUrlCreate(
	connFactory database.ConnFactory,
	query content.Query,
	repos content.UrlCreateRepository,
	service content.UrlService,
) content.UrlCreateUseCase {
	return &urlCreate{
		connFactory: connFactory,
		query:       query,
		repos:       repos,
		service:     service,
	}
}

type urlCreate struct {
	connFactory database.ConnFactory
	query       content.Query
	repos       content.UrlCreateRepository
	service     content.UrlService
}

func (uc *urlCreate) Exec(
	ctx context.Context,
	courseId vo.CourseId,
	in content.UrlInput,
) (*content.CreateOutput, error) {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[urlCreate.Exec]connFactory.Createのエラー", err)
	}

	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return nil, errs.Wrap("[urlCreate.Exec]ctxt.AuthenticatedUserのエラー", err)
	}

	exist, err := uc.query.ExistCourse(ctx, conn, courseId)
	if err != nil {
		return nil, errs.Wrap("[urlCreate.Exec]query.ExistCourseのエラー", err)
	} else if !exist {
		return nil, errs.NewNotFound("講習が存在しません")
	}

	validUrl, err := uc.service.NewValidatedUrl(in)
	if err != nil {
		return nil, errs.Wrap("[urlCreate.Exec]service.NewValidatedUrlのエラー", err)
	}

	newContent, err := uc.repos.Create(ctx, conn, authedUser, courseId, validUrl)
	if err != nil {
		return nil, errs.Wrap("[urlCreate.Exec]repos.Createのエラー", err)
	}

	return &content.CreateOutput{
		ContentId: newContent.Id,
	}, nil
}
