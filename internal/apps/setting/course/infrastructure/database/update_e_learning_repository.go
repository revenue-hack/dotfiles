package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gorm.io/gorm"
)

func NewUpdateELearningRepository() course.UpdateELearningRepository {
	return &updateELearningRepository{}
}

type updateELearningRepository struct{}

func (r *updateELearningRepository) Update(
	ctx context.Context,
	conn *database.Conn,
	authedUser *authed.User,
	courseId vo.CourseId,
	in model.ValidatedCourse,
) error {
	schedule, err := r.getSchedule(ctx, conn, authedUser, courseId, in)
	if err != nil {
		return errs.Wrap("[updateELearningRepository.Update]getScheduleのエラー", err)
	}

	return conn.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Updates(entity.Course{
			Id:                 courseId.Value(),
			Title:              in.Title,
			Description:        in.Description,
			ThumbnailImageName: in.ThumbnailImage,
			IsRequired:         in.IsRequired,
			CategoryId:         in.CategoryId,
			UpdatedBy:          authedUser.UserId(),
		}).Error; err != nil {
			return errs.Wrap("[updateELearningRepository.Update]coursesの更新エラー", err)
		}

		if err := tx.Save(schedule).Error; err != nil {
			return errs.Wrap("[updateELearningRepository.Update]course_schedules, e_learning_schedulesの更新エラー", err)
		}
		return nil
	})
}

func (r *updateELearningRepository) getSchedule(
	ctx context.Context,
	conn *database.Conn,
	authedUser *authed.User,
	courseId vo.CourseId,
	in model.ValidatedCourse,
) (*entity.CourseSchedule, error) {
	query := conn.DB().
		Preload("ELearningSchedule").
		Where("course_schedules.course_id = ?", courseId.Value())
	record, err := database.Get[entity.CourseSchedule](ctx, query)

	if err != nil {
		if database.IsErrRecordNotFound(err) {
			// レコードない場合はエラーではなく正常系として扱うので、登録用のデータを設定して値を返します
			return &entity.CourseSchedule{
				CreatedBy: authedUser.UserId(),
				UpdatedBy: authedUser.UserId(),
				ELearningSchedule: entity.ELearningSchedule{
					From:      in.From,
					To:        in.To,
					CreatedBy: authedUser.UserId(),
					UpdatedBy: authedUser.UserId(),
				},
			}, nil
		}
		return nil, errs.Wrap("[updateELearningRepository.getSchedule]course_schedulesの取得エラー", err)
	}

	// course_schedulesが存在する場合、必ずe_learning_schedulesは存在する
	record.UpdatedBy = authedUser.UserId()
	record.ELearningSchedule.From = in.From
	record.ELearningSchedule.To = in.To
	record.ELearningSchedule.UpdatedBy = authedUser.UserId()
	return record, nil
}
