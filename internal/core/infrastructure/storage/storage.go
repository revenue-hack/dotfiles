package storage

import (
	"context"
	"fmt"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

const (
	// サムネイル画像を配置するパスのフォーマット
	// {顧客コード}/{講習ID}/{画像名}
	thumbnailImagePath = "%s/%d/%s"
)

func NewFactory() Factory {
	return &factory{}
}

type Factory interface {
	Create(context.Context) (Client, error)
}

type Client interface {
	// Create は指定のファイルを作成します
	Create(ctx context.Context, distPath string, content []byte) error
}

// MakeThumbnailImagePath サムネイル画像を配置するためのパスを返却します
func MakeThumbnailImagePath(customerCode string, courseId vo.CourseId, imageName string) string {
	return fmt.Sprintf(thumbnailImagePath, customerCode, courseId.Value(), imageName)
}