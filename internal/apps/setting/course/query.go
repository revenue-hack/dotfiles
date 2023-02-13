package course

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
)

type GetQuery interface {
	// GetELearning はe-Learningの概要情報を取得します
	GetELearning(context.Context, *database.Conn, vo.CourseId) (*entity.Course, error)
}
