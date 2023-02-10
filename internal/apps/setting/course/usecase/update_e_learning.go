package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewUpdateELearning(
	connFactory database.ConnFactory,
	query course.UpdateQuery,
	repos course.UpdateELearningRepository,
	service course.UpdateELearningService,
) course.UpdateELearningUseCase {
	return &updateELearning{
		connFactory: connFactory,
		query:       query,
		repos:       repos,
		service:     service,
	}
}

type updateELearning struct {
	connFactory database.ConnFactory
	query       course.UpdateQuery
	repos       course.UpdateELearningRepository
	service     course.UpdateELearningService
}

func (uc *updateELearning) Exec(ctx context.Context, courseId vo.CourseId, in course.UpdateELearningInput) error {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return errs.Wrap("[updateELearning.Exec]connFactory.Createのエラー", err)
	}

	if exist, err := uc.query.ExistCourse(ctx, conn, courseId); err != nil {
		return errs.Wrap("[updateELearning.Exec]query.ExistCourseのエラー", err)
	} else if !exist {
		return errs.NewNotFound("講習が存在しません")
	}

	return nil
}
