package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

func NewSearchOptional(
	connFactory database.ConnFactory,
	query search.Query,
) search.UseCase {
	return &searchOptional{connFactory: connFactory, query: query}
}

type searchOptional struct {
	query       search.Query
	connFactory database.ConnFactory
}

func (uc *searchOptional) Exec(ctx context.Context, in search.Input) (*search.Output, error) {
	return &search.Output{}, nil
}
