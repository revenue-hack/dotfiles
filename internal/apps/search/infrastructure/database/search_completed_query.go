package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	ae "gitlab.kaonavi.jp/ae/sardine/internal/apps/search/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model/searchparam"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

func NewSearchCompletedQuery() search.Query {
	return &searchCompletedQuery{}
}

type searchCompletedQuery struct {
}

func (*searchCompletedQuery) GetMaxPageCount(
	ctx context.Context,
	conn *database.Conn,
	param searchparam.SearchParam,
) (uint32, error) {
	return 0, nil
}

func (*searchCompletedQuery) Get(
	ctx context.Context,
	conn *database.Conn,
	param searchparam.SearchParam,
) ([]ae.Course, error) {
	return nil, nil
}
