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
	schedule, elSchedule, err := r.getSchedule(ctx, conn, authedUser, courseId, in)
	if err != nil {
		return errs.Wrap("[updateELearningRepository.Update]getScheduleのエラー", err)
	}

	courseParams := map[string]any{
		"title":       in.Title,
		"description": in.Description,
		"is_required": in.IsRequired,
		"category_id": in.CategoryId,
		"updated_by":  authedUser.UserId(),
	}
	if in.Thumbnail != nil {
		courseParams["thumbnail_delivery_file_name"] = in.Thumbnail.Name
		courseParams["thumbnail_original_file_name"] = in.Thumbnail.OriginalName
	}

	return conn.Transaction(ctx, func(tx *gorm.DB) error {
		// 構造体を使うとis_required=falseがゼロ値で無視されてしまうのでmapを使用しています
		// TODO: サムネイル情報を更新する
		if err := tx.Model(&entity.Course{Id: courseId.Value()}).Updates(courseParams).Error; err != nil {
			return errs.Wrap("[updateELearningRepository.Update]coursesの更新エラー", err)
		}

		if err := tx.Save(schedule).Error; err != nil {
			return errs.Wrap("[updateELearningRepository.Update]course_schedulesの更新エラー", err)
		}

		// 更新の場合にIDがないとエラーになるので常に上書きしておく
		elSchedule.CourseScheduleId = schedule.Id
		if err := tx.Save(elSchedule).Error; err != nil {
			return errs.Wrap("[updateELearningRepository.Update]e_learning_schedulesの更新エラー", err)
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
) (*entity.CourseSchedule, *entity.ELearningSchedule, error) {
	query := conn.DB().
		Preload("ELearningSchedule").
		Where("course_schedules.course_id = ?", courseId.Value())
	record, err := database.Get[entity.CourseSchedule](ctx, query)

	if err == nil {
		record.UpdatedBy = authedUser.UserId()

		// course_schedulesが存在する場合、必ずe_learning_schedulesは存在する
		record.ELearningSchedule.From = in.From
		record.ELearningSchedule.To = in.To
		record.ELearningSchedule.UpdatedBy = authedUser.UserId()
		return record, &record.ELearningSchedule, nil
	}

	if !database.IsErrRecordNotFound(err) {
		return nil, nil, errs.Wrap("[updateELearningRepository.getSchedule]course_schedulesの取得エラー", err)
	}

	// レコードない場合はエラーではなく正常系として扱うので、登録用のデータを設定して値を返します
	return &entity.CourseSchedule{
			CourseId:  courseId.Value(),
			CreatedBy: authedUser.UserId(),
			UpdatedBy: authedUser.UserId(),
		},
		&entity.ELearningSchedule{
			From:      in.From,
			To:        in.To,
			CreatedBy: authedUser.UserId(),
			UpdatedBy: authedUser.UserId(),
		},
		nil
}
