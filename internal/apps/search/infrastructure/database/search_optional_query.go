package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model/searchparam"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

func NewSearchOptionalQuery() search.Query {
	return &searchOptionalQuery{}
}

type searchOptionalQuery struct {
}

func (h *searchOptionalQuery) GetMaxPageCount(
	ctx context.Context,
	conn *database.Conn,
	param searchparam.SearchParam,
) (uint32, error) {
	return 0, nil
}

func (h *searchOptionalQuery) Get(
	ctx context.Context,
	conn *database.Conn,
	param searchparam.SearchParam,
) ([]model.Course, error) {
	return nil, nil
}