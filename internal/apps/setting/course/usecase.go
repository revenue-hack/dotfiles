package course

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

type GetELearningUseCase interface {
	Exec(context.Context, vo.CourseId) (*GetELearningOutput, error)
}
