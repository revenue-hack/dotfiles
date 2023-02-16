package content

import (
	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
)

type ListOutput struct {
	Contents entity.Contents
}

type CreateOutput struct {
	ContentId uint32
}
