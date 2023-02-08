package search

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model/searchparam"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

type Query interface {
	// GetMaxPageCount は指定条件での最大ページ数を返却します
	GetMaxPageCount(context.Context, *database.Conn, searchparam.SearchParam) (uint32, error)
	// Get は指定条件の講習を全て取得します
	Get(context.Context, *database.Conn, searchparam.SearchParam) ([]model.Course, error)
}
