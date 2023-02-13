package vo

import (
	"strconv"

	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewCourseIdFromString(s string) (*CourseId, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, errs.NewInvalidParameter("講習IDは数値で指定してください")
	}
	return &CourseId{value: uint32(i)}, nil
}

type CourseId struct {
	value uint32
}

func (v CourseId) Value() uint32 {
	return v.value
}
