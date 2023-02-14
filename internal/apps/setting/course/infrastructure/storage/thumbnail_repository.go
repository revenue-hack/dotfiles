package storage

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	dc "gitlab.kaonavi.jp/ae/sardine/internal/core/domain/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/storage"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewThumbnailRepository(factory storage.Factory) course.ThumbnailRepository {
	return &thumbnailRepository{factory: factory}
}

type thumbnailRepository struct {
	factory storage.Factory
}

func (h *thumbnailRepository) Create(
	ctx context.Context,
	authedUser *authed.User,
	courseId vo.CourseId,
	thumb *model.ThumbnailImage,
) error {
	client, err := h.factory.Create(ctx)
	if err != nil {
		return errs.Wrap("[thumbnailRepository.Create]factory.Clientのエラー", err)
	}
	path := dc.MakeThumbnailImagePath(authedUser.CustomerCode(), courseId, "")
	return client.Create(ctx, path, thumb.Content)
}
