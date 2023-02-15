//go:build local || feature_test

package cdn

import (
	"context"
	"fmt"
	"strings"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
)

func NewFactory() Factory {
	return &factory{}
}

type factory struct{}

func (*factory) Create(ctx context.Context) (Client, error) {
	return &mockClient{
		baseUrl: env.GetString("AWS_CDN_BASE_URL"),
	}, nil
}

type mockClient struct {
	baseUrl string
}

func (c *mockClient) CreateUrl(path string) (string, error) {
	return fmt.Sprintf("%s/%s", c.baseUrl, strings.TrimPrefix(path, "/")), nil
}
