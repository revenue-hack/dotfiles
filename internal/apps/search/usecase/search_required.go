package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

func NewSearchRequired(
	connFactory database.ConnFactory,
	query search.Query,
) search.UseCase {
	return &searchRequired{connFactory: connFactory, query: query}
}

type searchRequired struct {
	query       search.Query
	connFactory database.ConnFactory
}

func (h *searchRequired) Exec(ctx context.Context, in search.Input) (*search.Output, error) {
	return &search.Output{}, nil
}
