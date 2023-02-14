package course

import (
	"fmt"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

const (
	// サムネイル画像を配置するパスのフォーマット
	// {顧客コード}/{講習ID}/{画像名}
	thumbnailImagePath = "%s/%d/%s"
	// 動画コンテンツを配置する基底ディレクトリのフォーマット
	// {顧客コード}/{講習ID}
	movieContentBasePath = "%s/%d/movies"
	// ファイルコンテンツを配置するパスのフォーマット
	// {顧客コード}/{講習ID}/{ファイル名}
	fileContentPath = "%s/%d/files/%s"
)

func MakeThumbnailImagePath(customerCode string, courseId vo.CourseId, imageName string) string {
	return fmt.Sprintf(thumbnailImagePath, customerCode, courseId.Value(), imageName)
}

func MakeMovieContentBasePath(customerCode string, courseId vo.CourseId) string {
	return fmt.Sprintf(movieContentBasePath, customerCode, courseId.Value())
}

func MakeFileContentPath(customerCode string, courseId vo.CourseId, imageName string) string {
	return fmt.Sprintf(fileContentPath, customerCode, courseId.Value(), imageName)
}
