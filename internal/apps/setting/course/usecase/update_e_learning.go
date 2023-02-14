package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewUpdateELearning(
	connFactory database.ConnFactory,
	query course.UpdateQuery,
	repos course.UpdateELearningRepository,
	thumbRepos course.ThumbnailRepository,
	service course.UpdateELearningService,
) course.UpdateELearningUseCase {
	return &updateELearning{
		connFactory: connFactory,
		query:       query,
		repos:       repos,
		thumbRepos:  thumbRepos,
		service:     service,
	}
}

type updateELearning struct {
	connFactory database.ConnFactory
	query       course.UpdateQuery
	repos       course.UpdateELearningRepository
	thumbRepos  course.ThumbnailRepository
	service     course.UpdateELearningService
}

func (uc *updateELearning) Exec(ctx context.Context, courseId vo.CourseId, in course.UpdateELearningInput) error {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return errs.Wrap("[updateELearning.Exec]connFactory.Createのエラー", err)
	}

	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return errs.Wrap("[updateELearning.Exec]ctxt.AuthenticatedUserのエラー", err)
	}

	exist, err := uc.query.ExistCourse(ctx, conn, courseId)
	if err != nil {
		return errs.Wrap("[updateELearning.Exec]query.ExistCourseのエラー", err)
	} else if !exist {
		return errs.NewNotFound("講習が存在しません")
	}

	valid, err := uc.service.NewValidatedCourse(ctx, conn, in)
	if err != nil {
		return errs.Wrap("[updateELearning.Exec]service.NewValidatedCourseのエラー", err)
	}

	if err = uc.repos.Update(ctx, conn, authedUser, courseId, *valid); err != nil {
		return errs.Wrap("[updateELearning.Exec]repos.Updateのエラー", err)
	}

	// サムネイル画像がある場合だけ生成
	if valid.ThumbnailImage != nil {
		if err = uc.thumbRepos.Create(ctx, authedUser, courseId, valid.ThumbnailImage); err != nil {
			return errs.Wrap("[updateELearning.Exec]thumbRepos.Createのエラー", err)
		}
	}

	return nil
}
