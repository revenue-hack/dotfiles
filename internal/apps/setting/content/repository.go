package content

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
)

type UrlCreateRepository interface {
	// Create は外部URLコンテンツを登録します
	Create(context.Context, *database.Conn, *authed.User, vo.CourseId, *model.ValidatedUrl) (*entity.Content, error)
}

type UrlUpdateRepository interface {
	// Update は外部URLコンテンツを更新します
	Update(context.Context, *database.Conn, *authed.User, vo.ContentId, *model.ValidatedUrl) error
}

type UrlDeleteRepository interface {
	// Delete は外部URLコンテンツを削除します
	Delete(context.Context, *database.Conn, vo.ContentId) error
}
