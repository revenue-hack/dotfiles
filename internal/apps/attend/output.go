package attend

import (
	"time"

	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
)

type GetELearningOutput struct {
	Title       string
	Description *string
	From        *time.Time
	To          *time.Time
	IsFixed     bool
	Contents    entity.Contents
}
