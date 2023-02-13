package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	ae "gitlab.kaonavi.jp/ae/sardine/internal/apps/search/entity"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model/searchparam"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewSearchRequiredQuery() search.Query {
	return &searchRequiredQuery{}
}

type searchRequiredQuery struct {
}

func (h *searchRequiredQuery) GetMaxPageCount(
	ctx context.Context,
	conn *database.Conn,
	param searchparam.SearchParam,
) (uint32, error) {
	return 0, nil
}

func (h *searchRequiredQuery) Get(
	ctx context.Context,
	conn *database.Conn,
	_ searchparam.SearchParam,
) ([]ae.Course, error) {
	query := conn.DB().
		Select([]string{
			"courses.id",
			"courses.title",
			"courses.thumbnail_image_name",
			"categories.name as category_name",
			"els.from",
			"els.to",
			"courses.is_required",
		}).
		Table("courses").
		Joins("LEFT JOIN categories ON categories.id = courses.category_id").
		Joins("LEFT JOIN course_schedules cs ON courses.id = cs.course_id").
		Joins("LEFT JOIN e_learning_schedules els ON cs.id = els.course_schedule_id")

	// TODO: 一旦入っているデータを全部返したいので非公開、必須の条件を外してかえす + あとでLimitとOrderを指定する
	// Where("courses.status = ?", entity.CourseStatusPublic)
	// Where("courses.is_required = ?", entity.CourseIsRequired)

	records, err := database.GetAll[ae.Course](ctx, query)
	if err != nil {
		return nil, errs.Wrap("[searchRequiredQuery.Get]database.GetAllのエラー", err)
	}
	return records, nil
}
