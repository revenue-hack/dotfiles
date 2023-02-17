package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/attend"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gorm.io/gorm"
)

func NewGetQuery() attend.GetQuery {
	return &getQuery{}
}

type getQuery struct{}

func (*getQuery) GetELearning(
	ctx context.Context,
	conn *database.Conn,
	courseId vo.CourseId,
) (*entity.Course, error) {
	query := conn.DB().
		Select([]string{
			"c.id",
			"c.title",
			"c.description",
		}).
		Table("courses c").
		Preload("CourseSchedules.ELearningSchedule").
		Preload("Contents", func(db *gorm.DB) *gorm.DB {
			return db.Order("contents.display_order ASC")
		}).
		Preload("Contents.Url").
		Where("c.id = ?", courseId.Value()).
		Where("c.course_type = ?", entity.CourseTypeELearning).
		Where("c.status = ?", entity.CourseStatusPublic)

	record, err := database.Get[entity.Course](ctx, query)
	if err != nil {
		return nil, errs.Wrap("[getQuery.GetELearning]database.Getのエラー", err)
	}
	return record, nil
}
