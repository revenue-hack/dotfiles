package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/attend"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/domain/elearning"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/cdn"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewGetELearning(
	connFactory database.ConnFactory,
	cdnFactory cdn.Factory,
	query attend.GetQuery,
) attend.GetELearningUseCase {
	return &getELearning{
		connFactory: connFactory,
		cdnFactory:  cdnFactory,
		query:       query,
	}
}

type getELearning struct {
	connFactory database.ConnFactory
	cdnFactory  cdn.Factory
	query       attend.GetQuery
}

func (uc *getELearning) Exec(ctx context.Context, courseId vo.CourseId) (*attend.GetELearningOutput, error) {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[getELearning.Exec]connFactory.Createのエラー", err)
	}

	record, err := uc.query.GetELearning(ctx, conn, courseId)
	if err != nil {
		if database.IsErrRecordNotFound(err) {
			return nil, errs.NewNotFound("講習が存在しません")
		}
		return nil, errs.Wrap("[getELearning.Exec]query.GetELearningのエラー", err)
	}

	return &attend.GetELearningOutput{
		Title:       record.Title,
		Description: record.Description,
		From:        elearning.GetFrom(record),
		To:          elearning.GetTo(record),
		IsFixed:     false, // TODO: 受講状況を見て値を返す
		Contents:    record.Contents,
	}, nil
}
