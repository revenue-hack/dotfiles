package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gorm.io/gorm"
)

func NewCreateRepository() course.CreateRepository {
	return &createRepository{}
}

type createRepository struct{}

func (*createRepository) CreateELearning(
	ctx context.Context,
	conn *database.Conn,
	authedUser *authed.User,
) (*entity.Course, error) {
	record := &entity.Course{
		Title:      "無題のe-Learning",
		CourseType: entity.CourseTypeELearning,
		IsRequired: false,
		Status:     entity.CourseStatusPrivate,
		CreatedBy:  authedUser.UserId(),
		UpdatedBy:  authedUser.UserId(),
	}
	err := conn.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Create(record).Error; err != nil {
			return errs.Wrap("[createRepository.CreateELearning]coursesの保存に失敗", err)
		}
		return nil
	})
	if err != nil {
		return nil, errs.Wrap("[createRepository.CreateELearning]DBの保存に失敗", err)
	}
	return record, nil
}
