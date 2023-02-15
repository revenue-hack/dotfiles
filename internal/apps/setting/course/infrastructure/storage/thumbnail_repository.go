package storage

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/domain/file"
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
	thumb *file.UploadFile,
) error {
	client, err := h.factory.Create(ctx)
	if err != nil {
		return errs.Wrap("[thumbnailRepository.Create]factory.Clientのエラー", err)
	}
	path := storage.MakeThumbnailImagePath(authedUser.CustomerCode(), courseId.Value(), thumb.HashedName())
	return client.Create(ctx, path, thumb.Content())
}
