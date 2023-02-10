package course

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
)

type CreateRepository interface {
	// CreateELearning はe-Learning講習を作成します
	CreateELearning(context.Context, *database.Conn, *authed.User) (*entity.Course, error)
}
