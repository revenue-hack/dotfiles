package ctxt

import (
	"context"
	"fmt"
)

// Contextのキーを表す型定義です
type key string

const (
	ContextKeyLogUniqueId key = "log_unique_id"
)

// RequestId リクエスト固有のIDを取得します
func RequestId(ctx context.Context) (string, error) {
	return get[string](ctx, ContextKeyLogUniqueId)
}

func get[T any](ctx context.Context, k key) (T, error) {
	v, ok := ctx.Value(k).(T)
	if !ok {
		// TODO: errsパッケージを作ったら置き換える
		return v, fmt.Errorf("Contextに%sが存在しません。", k)
	}
	return v, nil
}

func (r key) String() string {
	return string(r)
}
