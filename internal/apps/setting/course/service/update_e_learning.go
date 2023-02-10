package service

import (
	"context"
	"time"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/timer"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/validate"
)

func NewUpdateELearning() course.UpdateELearningService {
	return &updateELearning{}
}

type updateELearning struct{}

func (*updateELearning) NewValidatedCourse(
	ctx context.Context,
	conn *database.Conn,
	in course.UpdateELearningInput,
) (*model.ValidatedCourse, error) {
	ers := errs.NewErrors()
	ers.AddError(validate.StringRequired("講習タイトル", &in.Title, 50))
	ers.AddError(validate.StringOptional("講習の説明", in.Description, 1000))
	ers.AddError(validate.StringOptional("サムネイル画像", in.ThumbnailImage, 100))

	from, err := parseDatetime("期間（開始）", in.From)
	ers.AddError(err)
	to, err := parseDatetime("期間（終了）", in.To)
	ers.AddError(err)

	// 期間指定があれば to > from である必要がある
	if from != nil && to != nil && to.Before(*from) {
		ers.AddError(errs.NewInvalidParameter("期間（開始）は期間（終了）の過去日時を指定してください"))
	}

	// TODO: カテゴリの検証

	return errs.ErrorsOrNilWithValue(model.ValidatedCourse{
		Title:          in.Title,
		Description:    in.Description,
		ThumbnailImage: in.ThumbnailImage,
		IsRequired:     in.IsRequired,
		CategoryId:     in.CategoryId,
		From:           from,
		To:             to,
	}, ers)
}

// 日付をパースします
func parseDatetime(fieldName string, val *string) (*time.Time, error) {
	if val == nil || *val == "" {
		return nil, nil
	}
	t, err := timer.Parse("2006/01/02 15:04", *val)
	if err != nil {
		return nil, errs.NewInvalidParameter("%sの日付形式が不正です", fieldName)
	}
	return &t, nil
}
