package helper

// P は引数のポインタを返却します
func P[T any](v T) *T {
	return &v
}
