//go:build !local && !feature_test

package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

type factory struct{}

func (*factory) Create(ctx context.Context) (Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, errs.Wrap("[storage.NewClient]AWS接続用の設定読み込みエラー", err)
	}
	return &s3Client{
		client: s3.NewFromConfig(cfg),
		bucket: env.GetString("AWS_S3_DELIVERY_BUCKET"),
	}, nil
}
