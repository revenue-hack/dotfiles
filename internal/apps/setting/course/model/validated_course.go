package model

import (
	"time"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/timer"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/validate"
)

// ValidatedCourse は検証済みの講習情報です
type ValidatedCourse struct {
	title          string
	description    *string
	thumbnailImage *string
	isRequired     bool
	categoryId     *uint32
	from           *time.Time
	to             *time.Time
}

func (v ValidatedCourse) Title() string {
	return v.title
}

func (v ValidatedCourse) Description() *string {
	return v.description
}

func (v ValidatedCourse) ThumbnailImage() *string {
	return v.thumbnailImage
}

func (v ValidatedCourse) IsRequired() bool {
	return v.isRequired
}

func (v ValidatedCourse) CategoryId() *uint32 {
	return v.categoryId
}

func (v ValidatedCourse) From() *time.Time {
	return v.from
}

func (v ValidatedCourse) To() *time.Time {
	return v.to
}

func NewValidatedCourse(in course.UpdateELearningInput) (*ValidatedCourse, error) {
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

	return errs.ErrorsOrNilWithValue(&ValidatedCourse{
		title:          in.Title,
		description:    in.Description,
		thumbnailImage: in.ThumbnailImage,
		isRequired:     in.IsRequired,
		categoryId:     in.CategoryId,
		from:           from,
		to:             to,
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
