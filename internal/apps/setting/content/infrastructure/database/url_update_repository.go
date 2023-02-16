package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gorm.io/gorm"
)

func NewUrlUpdateRepository() content.UrlUpdateRepository {
	return &urlUpdateRepository{}
}

type urlUpdateRepository struct{}

func (r *urlUpdateRepository) Update(
	ctx context.Context,
	conn *database.Conn,
	authedUser *authed.User,
	contentId vo.ContentId,
	in *model.ValidatedUrl,
) error {
	newContent := &entity.Content{
		Id:        contentId.Value(),
		UpdatedBy: authedUser.UserId(),
	}
	newUrl := &entity.Url{
		Title:     in.Title,
		Url:       in.Url,
		UpdatedBy: authedUser.UserId(),
	}

	err := conn.Transaction(ctx, func(tx *gorm.DB) error {
		if txErr := tx.Updates(newContent).Error; txErr != nil {
			return errs.Wrap("[urlUpdateRepository.Update]contentsの更新に失敗", txErr)
		}
		if txErr := tx.Where("content_id = ?", contentId.Value()).Updates(newUrl).Error; txErr != nil {
			return errs.Wrap("[urlUpdateRepository.Update]urlsの更新に失敗", txErr)
		}
		return nil
	})

	if err != nil {
		return errs.Wrap("[urlUpdateRepository.Update]外部URLコンテンツの更新に失敗", err)
	}
	return nil
}
