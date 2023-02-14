package storage

import (
	"context"
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
