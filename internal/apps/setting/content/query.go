package content

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
)

type GetQuery interface {
	// GetAll は全種別のコンテンツ情報をまとめて取得します
	GetAll(context.Context, *database.Conn, vo.CourseId) (entity.Contents, error)
}
