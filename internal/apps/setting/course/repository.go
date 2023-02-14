package course

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

type UpdateELearningRepository interface {
	// Update はe-Learningの概要情報を更新します
	Update(context.Context, *database.Conn, *authed.User, vo.CourseId, model.ValidatedCourse) error
}

type ThumbnailRepository interface {
	// Create はサムネイル画像を生成します
	Create(context.Context, *authed.User, vo.CourseId, *model.ThumbnailImage) error
}
