package search

import (
	"context"
)

type UseCase interface {
	Exec(context.Context, Input) (*Output, error)
}
