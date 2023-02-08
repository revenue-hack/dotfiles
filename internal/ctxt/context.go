package ctxt

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

// Contextのキーを表す型定義です
type key string

const (
	KeyLogUniqueId       key = "log_unique_id"
	KeyAuthenticatedUser key = "authenticated_user"
)

// RequestId リクエスト固有のIDを取得します
func RequestId(ctx context.Context) (string, error) {
	return get[string](ctx, KeyLogUniqueId)
}

// AuthenticatedUser 認証済みのユーザー情報を取得します
func AuthenticatedUser(ctx context.Context) (*authed.User, error) {
	return get[*authed.User](ctx, KeyAuthenticatedUser)
}

func get[T any](ctx context.Context, k key) (T, error) {
	v, ok := ctx.Value(k).(T)
	if !ok {
		return v, errs.NewInternalError("contextに%sが存在しません", k)
	}
	return v, nil
}
