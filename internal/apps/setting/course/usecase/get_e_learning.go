package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/domain/elearning"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/cdn"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/storage"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewGetELearning(
	connFactory database.ConnFactory,
	cdnFactory cdn.Factory,
	query course.GetQuery,
) course.GetELearningUseCase {
	return &getELearning{
		connFactory: connFactory,
		cdnFactory:  cdnFactory,
		query:       query,
	}
}

type getELearning struct {
	connFactory database.ConnFactory
	cdnFactory  cdn.Factory
	query       course.GetQuery
}

func (uc *getELearning) Exec(ctx context.Context, courseId vo.CourseId) (*course.GetELearningOutput, error) {
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

	out, err := uc.makeOutput(ctx, courseId, record)
	if err != nil {
		return nil, errs.Wrap("[getELearning.Exec]makeOutputのエラー", err)
	}
	return out, nil
}

func (uc *getELearning) makeOutput(
	ctx context.Context,
	courseId vo.CourseId,
	c *entity.Course,
) (*course.GetELearningOutput, error) {
	out := &course.GetELearningOutput{
		Id:          c.Id,
		Title:       c.Title,
		Description: c.Description,
		IsRequired:  c.IsRequired,
		CategoryId:  c.CategoryId,
		From:        elearning.GetFrom(c),
		To:          elearning.GetTo(c),
	}

	if !c.HasThumbnail() {
		return out, nil
	}

	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return nil, errs.Wrap("[getELearning.bindOutput]ctxt.AuthenticatedUserのエラー", err)
	}

	// サムネイル画像がある場合は配信用のURLに変換する
	client, err := uc.cdnFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[getELearning.bindOutput]cdnFactory.Createのエラー", err)
	}
	// TODO: 画像パスをDBに保存したほうが楽かも
	url, err := client.CreateUrl(
		storage.MakeThumbnailImagePath(authedUser.CustomerCode(), courseId, *c.ThumbnailDeliveryFileName))
	if err != nil {
		return nil, errs.Wrap("[getELearning.bindOutput]client.CreateUrlのエラー", err)
	}

	out.Thumbnail = &course.GetELearningThumbnailOutput{
		Name: *c.ThumbnailOriginalFileName,
		Url:  url,
	}
	return out, nil
}
