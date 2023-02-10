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

func (r *createRepository) CreateELearning(
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
		return tx.Create(record).Error
	})
	if err != nil {
		return nil, errs.NewInternalError("failed to createRepository.CreateELearning: %v", err)
	}
	return record, nil
}
