package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gorm.io/gorm"
)

func NewUrlDeleteRepository() content.UrlDeleteRepository {
	return &urlDeleteRepository{}
}

type urlDeleteRepository struct{}

func (r *urlDeleteRepository) Delete(ctx context.Context, conn *database.Conn, contentId vo.ContentId) error {
	err := conn.Transaction(ctx, func(tx *gorm.DB) error {
		if txErr := tx.Delete(&entity.Content{Id: contentId.Value()}).Error; txErr != nil {
			return errs.Wrap("[urlUpdateRepository.Update]contentsの削除に失敗", txErr)
		}
		return nil
	})

	if err != nil {
		return errs.Wrap("[urlUpdateRepository.Update]外部URLコンテンツの削除に失敗", err)
	}
	return nil
}
