//go:build feature_test

// feature test用に値を変更したい定数を定義してください

package file

import (
	_ "image/gif"  // image.Decodeに必要
	_ "image/jpeg" // image.Decodeに必要
	_ "image/png"  // image.Decodeに必要
)

const (
	// 画像の最大サイズ(1MB)
	imageSizeLimit = 1 << 20
)
