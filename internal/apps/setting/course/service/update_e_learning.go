package service

import (
	"context"
	"encoding/base64"
	"net/http"
	"path/filepath"
	"time"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/hash"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/timer"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/validate"
)

var (
	validMimeTypes = map[string]struct{}{
		"image/jpeg": {},
		"image/png":  {},
		"image/gif":  {},
	}
)

func NewUpdateELearning(query course.UpdateQuery) course.UpdateELearningService {
	return &updateELearning{query: query}
}

type updateELearning struct {
	query course.UpdateQuery
}

func (s *updateELearning) NewValidatedCourse(
	ctx context.Context,
	conn *database.Conn,
	in course.UpdateELearningInput,
) (*model.ValidatedCourse, error) {
	ers := errs.NewErrors()
	ers.AddError(validate.StringRequired("講習タイトル", &in.Title, 50))
	ers.AddError(validate.StringOptional("講習の説明", in.Description, 1000))

	from, err := s.parseDatetime("期間（開始）", in.From)
	ers.AddError(err)
	to, err := s.parseDatetime("期間（終了）", in.To)
	ers.AddError(err)

	// 期間指定があれば to > from である必要がある
	if from != nil && to != nil && (to.Equal(*from) || to.Before(*from)) {
		ers.Add("期間（開始）は期間（終了）の過去日時を指定してください")
	}

	if in.CategoryId != nil {
		exist, err := s.query.ExistCategory(ctx, conn, *in.CategoryId)
		if err != nil {
			return nil, errs.Wrap("[updateELearning.NewValidatedCourse]query.ExistCategoryのエラー", err)
		}
		if !exist {
			ers.Add("カテゴリが存在しません")
		}
	}

	// サムネイル画像
	thumb, thumbErs := s.parseThumbnail(in.Thumbnail)
	ers.Append(thumbErs)

	return errs.ErrorsOrNilWithValue(model.ValidatedCourse{
		Title:             in.Title,
		Description:       in.Description,
		ThumbnailImage:    thumb,
		IsRemoveThumbnail: in.IsRemoveThumbnailImage,
		IsRequired:        in.IsRequired,
		CategoryId:        in.CategoryId,
		From:              from,
		To:                to,
	}, ers)
}

// 日付をパースします
func (*updateELearning) parseDatetime(fieldName string, val *string) (*time.Time, error) {
	if val == nil || *val == "" {
		return nil, nil
	}
	t, err := timer.Parse("2006/01/02 15:04", *val)
	if err != nil {
		return nil, errs.NewInvalidParameter("%sの日付形式が不正です", fieldName)
	}
	return &t, nil
}

func (*updateELearning) parseThumbnail(in *course.UpdateThumbnailInput) (*model.ThumbnailImage, *errs.Errors) {
	if in == nil {
		return nil, nil
	}

	ers := errs.NewErrors()

	// 念の為ファイル名だけにしておく
	name := filepath.Base(in.Name)
	ers.AddError(validate.StringRequired("サムネイル画像名", &name, 50))

	// TODO: 拡張子の検証

	// コンテンツをDecodeして検証します
	decodedContent, err := base64.StdEncoding.DecodeString(in.Content)
	if err != nil {
		ers.Add("サムネイル画像が不正な形式です")
		return nil, ers
	}

	// MIME-Typeを検証
	if _, ok := validMimeTypes[http.DetectContentType(decodedContent)]; !ok {
		ers.Add("サムネイル画像はjpeg/png/gifを指定してください")
	}

	if ers.HasError() {
		return nil, ers
	}

	return &model.ThumbnailImage{
		OriginalName: name,
		Name:         hash.Make(name), // 元のファイル名を元にハッシュ化
		Content:      decodedContent,
	}, nil
}
