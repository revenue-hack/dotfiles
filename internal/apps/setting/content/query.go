package content

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
)

type Query interface {
	// ExistCourse は講習が存在する場合にtrueを返却します
	ExistCourse(context.Context, *database.Conn, vo.CourseId) (bool, error)
	// ExistUrl は外部URLコンテンツが存在する場合にtrueを返却します
	ExistUrl(context.Context, *database.Conn, vo.CourseId, vo.ContentId) (bool, error)
}

type ListQuery interface {
	// GetAll は全種別のコンテンツ情報をまとめて取得します
	GetAll(context.Context, *database.Conn, vo.CourseId) (entity.Contents, error)
}
