package course

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

type UpdateELearningService interface {
	// NewValidatedCourse はinputを検証して、検証済みのmodelを返却します
	NewValidatedCourse(context.Context, *database.Conn, UpdateELearningInput) (*model.ValidatedCourse, error)
}
