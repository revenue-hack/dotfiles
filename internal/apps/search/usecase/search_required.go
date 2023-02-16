package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/cdn"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/storage"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewSearchRequired(
	connFactory database.ConnFactory,
	cdnFactory cdn.Factory,
	query search.Query,
) search.UseCase {
	return &searchRequired{
		connFactory: connFactory,
		cdnFactory:  cdnFactory,
		query:       query,
	}
}

type searchRequired struct {
	query       search.Query
	connFactory database.ConnFactory
	cdnFactory  cdn.Factory
}

func (uc *searchRequired) Exec(ctx context.Context, in search.Input) (*search.Output, error) {
	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return nil, errs.Wrap("[searchRequired.Exec]ctxt.AuthenticatedUserのエラー", err)
	}

	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[searchRequired.Exec]connFactory.Createのエラー", err)
	}

	cdnClient, err := uc.cdnFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[searchRequired.Exec]cdnFactory.Createのエラー", err)
	}

	courses, err := uc.query.Get(ctx, conn, nil)
	if err != nil {
		return nil, errs.Wrap("[searchRequired.Exec]query.Getのエラー", err)
	}

	// modelにつめかえ
	ret := make([]model.Course, 0, len(courses))
	for _, c := range courses {
		mc := model.Course{
			Id:         c.Id,
			Title:      c.Title,
			ExpireAt:   c.To,
			IsRequired: c.IsRequired,
			IsFixed:    false, // 期限内のデータしかないはずなので固定値
		}

		// TODO: サムネ有無の判定はここ以外のどこかでやりたいが、とりあえず他の検索API作る時に考える
		if c.ThumbnailOriginalFileName != nil && c.ThumbnailDeliveryFileName != nil {
			// TODO: 画像パスをDBに保存したほうが楽かも
			url, err := cdnClient.CreateUrl(
				storage.MakeThumbnailImagePath(authedUser.CustomerCode(), c.Id, *c.ThumbnailDeliveryFileName))
			if err != nil {
				return nil, errs.Wrap("[searchRequired.Exec]cdnClient.CreateUrlのエラー", err)
			}
			mc.Thumbnail = &model.Thumbnail{
				Name: *c.ThumbnailOriginalFileName,
				Url:  url,
			}
		}

		ret = append(ret, mc)
	}

	return &search.Output{
		Courses: ret,
	}, nil
}
