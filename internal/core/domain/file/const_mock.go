//go:build feature_test

// feature test用に値を変更したい定数を定義してください

package file

const (
	// 画像の最大サイズ(1MB)
	imageSizeLimit = 1 << 20
)
