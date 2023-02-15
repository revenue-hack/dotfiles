//go:build !feature_test

package file

import (
	_ "image/gif"  // image.Decodeに必要
	_ "image/jpeg" // image.Decodeに必要
	_ "image/png"  // image.Decodeに必要
)

const (
	// 画像の最大サイズ(25MB)
	imageSizeLimit = 25 << 20
)
