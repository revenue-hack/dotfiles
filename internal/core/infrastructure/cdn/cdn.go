package cdn

import (
	"context"
)

type Factory interface {
	Create(context.Context) (Client, error)
}

type Client interface {
	// CreateUrl は配信用のURLを生成して返却します
	CreateUrl(string) (string, error)
}
