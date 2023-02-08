package search

import (
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/model"
)

type Output struct {
	MaxPageSize   uint32
	NextPageToken *string
	Courses       []model.Course
}
