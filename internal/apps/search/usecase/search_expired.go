package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

func NewSearchExpired(
	connFactory database.ConnFactory,
	query search.Query,
) search.UseCase {
	return &searchExpired{connFactory: connFactory, query: query}
}

type searchExpired struct {
	query       search.Query
	connFactory database.ConnFactory
}

func (h *searchExpired) Exec(ctx context.Context, in search.Input) (*search.Output, error) {
	return &search.Output{}, nil
}
