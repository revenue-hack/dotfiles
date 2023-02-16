package content

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

type ListUseCase interface {
	Exec(context.Context, vo.CourseId) (*ListOutput, error)
}
