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
) course.UpdateELearningUseCase {
	return &updateELearning{
		connFactory: connFactory,
		query:       query,
		repos:       repos,
	}
}

type updateELearning struct {
	connFactory database.ConnFactory
	query       course.UpdateQuery
	repos       course.UpdateELearningRepository
}

func (uc *updateELearning) Exec(ctx context.Context, courseId vo.CourseId, in course.UpdateELearningInput) error {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return errs.NewInternalError("failed to connFactory.Create from updateELearning: %v", err)
	}

	if exist, err := uc.query.ExistCourse(ctx, conn, courseId); err != nil {
		return errs.NewInternalError("failed to GetELearning from updateELearning: %v", err)
	} else if !exist {
		return errs.NewNotFound("講習が存在しません")
	}

	return nil
}
