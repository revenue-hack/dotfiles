package course

import (
	"context"
)

type CreateUseCase interface {
	Exec(context.Context) (*CreateOutput, error)
}
