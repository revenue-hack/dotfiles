//go:build feature_test

// Feature Test専用のClientを生成します
// Feature TestではS3を使用しないため

package storage

import (
	"context"
	"fmt"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
)

type factory struct{}

func (*factory) Create(ctx context.Context) (Client, error) {
	return &mockClient{
		bucket: fmt.Sprintf("%s/storage", env.GetString("TEST_BASE_DIR")),
	}, nil
}
