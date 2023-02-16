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
