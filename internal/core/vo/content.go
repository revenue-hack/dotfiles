package vo

import (
	"strconv"

	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

func NewContentIdFromString(s string) (*ContentId, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, errs.NewInvalidParameter("コンテンツIDは数値で指定してください")
	}
	return &ContentId{value: uint32(i)}, nil
}

type ContentId struct {
	value uint32
}

func (v ContentId) Value() uint32 {
	return v.value
}
