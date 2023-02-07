package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

func NewSearchCompleted(
	connFactory database.ConnFactory,
	query search.Query,
) search.UseCase {
	return &searchCompleted{connFactory: connFactory, query: query}
}

type searchCompleted struct {
	query       search.Query
	connFactory database.ConnFactory
}

func (h *searchCompleted) Exec(ctx context.Context, in search.Input) (*search.Output, error) {
	return &search.Output{}, nil
}
