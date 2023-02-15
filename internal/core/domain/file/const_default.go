//go:build !feature_test

package file

const (
	// 画像の最大サイズ(25MB)
	imageSizeLimit = 25 << 20
)
