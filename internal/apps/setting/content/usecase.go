package content

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

type ListUseCase interface {
	Exec(context.Context, vo.CourseId) (*ListOutput, error)
}

type UrlCreateUseCase interface {
	Exec(context.Context, vo.CourseId, UrlInput) (*CreateOutput, error)
}

type UrlUpdateUseCase interface {
	Exec(context.Context, vo.CourseId, vo.ContentId, UrlInput) error
}

type UrlDeleteUseCase interface {
	Exec(context.Context, vo.CourseId, vo.ContentId) error
}
