package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewGetELearning(
	connFactory database.ConnFactory,
	query course.GetQuery,
) course.GetELearningUseCase {
	return &getELearning{connFactory: connFactory, query: query}
}

type getELearning struct {
	connFactory database.ConnFactory
	query       course.GetQuery
}

func (uc *getELearning) Exec(ctx context.Context, courseId vo.CourseId) (*course.GetELearningOutput, error) {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[getELearning.Exec]connFactory.Createのエラー", err)
	}

	record, err := uc.query.GetELearning(ctx, conn, courseId)
	if err != nil {
		if err == database.ErrRecordNotFound {
			return nil, errs.NewNotFound("講習が存在しません")
		}
		return nil, errs.Wrap("[getELearning.Exec]query.GetELearningのエラー", err)
	}

	var out course.GetELearningOutput
	uc.bindOutput(record, &out)
	return &out, nil
}

func (uc *getELearning) bindOutput(c *entity.Course, out *course.GetELearningOutput) {
	out.Id = c.Id
	out.Title = c.Title
	out.Description = c.Description
	out.IsRequired = c.IsRequired
	out.CategoryId = c.CategoryId
	// TODO: サムネイル画像がある場合は配信用のURLに変換する
	out.ThumbnailUrl = c.ThumbnailImageName

	// 実施期間の指定がある場合のみFrom/Toをバインド
	if len(c.CourseSchedules) == 0 {
		return
	}

	// e-Learningは期間が1件しかないはずなので先頭データを使う
	out.From = c.CourseSchedules[0].ELearningSchedule.From
	out.To = c.CourseSchedules[0].ELearningSchedule.To
}
