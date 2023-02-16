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

func NewUrlCreateRepository() content.UrlCreateRepository {
	return &urlCreateRepository{}
}

type urlCreateRepository struct{}

func (r *urlCreateRepository) Create(
	ctx context.Context,
	conn *database.Conn,
	authedUser *authed.User,
	courseId vo.CourseId,
	in *model.ValidatedUrl,
) (*entity.Content, error) {
	var displayOrder uint16
	query := conn.DB().
		Model(&entity.Content{}).
		Select("COALESCE(MAX(display_order), 0)").
		Where("course_id = ?", courseId.Value())
	if err := query.First(&displayOrder).Error; err != nil {
		return nil, errs.Wrap("[urlCreateRepository.Create]display_orderの取得に失敗", err)
	}

	newContent := &entity.Content{
		CourseId:     courseId.Value(),
		ContentType:  entity.ContentTypeUrl,
		DisplayOrder: displayOrder + 1,
		CreatedBy:    authedUser.UserId(),
		UpdatedBy:    authedUser.UserId(),
	}
	newUrl := &entity.Url{
		Title:     in.Title,
		Url:       in.Url,
		CreatedBy: authedUser.UserId(),
		UpdatedBy: authedUser.UserId(),
	}

	err := conn.Transaction(ctx, func(tx *gorm.DB) error {
		if txErr := tx.Create(newContent).Error; txErr != nil {
			return errs.Wrap("[urlCreateRepository.Create]contentsの登録に失敗", txErr)
		}

		// contentの登録後にIDが払い出される
		newUrl.ContentId = newContent.Id
		if txErr := tx.Create(newUrl).Error; txErr != nil {
			return errs.Wrap("[urlCreateRepository.Create]urlsの登録に失敗", txErr)
		}
		return nil
	})

	if err != nil {
		return nil, errs.Wrap("[urlCreateRepository.Create]外部URLコンテンツの登録に失敗", err)
	}
	return newContent, nil
}
