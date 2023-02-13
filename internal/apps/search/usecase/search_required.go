package usecase

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewSearchRequired(
	connFactory database.ConnFactory,
	query search.Query,
) search.UseCase {
	return &searchRequired{connFactory: connFactory, query: query}
}

type searchRequired struct {
	query       search.Query
	connFactory database.ConnFactory
}

func (uc *searchRequired) Exec(ctx context.Context, in search.Input) (*search.Output, error) {
	conn, err := uc.connFactory.Create(ctx)
	if err != nil {
		return nil, errs.Wrap("[searchRequired.Exec]connFactory.Createのエラー", err)
	}

	courses, err := uc.query.Get(ctx, conn, nil)
	if err != nil {
		return nil, errs.Wrap("[searchRequired.Exec]query.Getのエラー", err)
	}

	// modelにつめかえ
	ret := make([]model.Course, 0, len(courses))
	for _, c := range courses {
		ret = append(ret, model.Course{
			Id:         c.Id,
			Title:      c.Title,
			ExpireAt:   c.To,
			IsRequired: c.IsRequired,
			IsFixed:    false, // 期限内のデータしかないはずなので固定値
		})
	}

	return &search.Output{
		Courses: ret,
	}, nil
}
