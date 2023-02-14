//go:build local

// ローカル専用のs3.Clientを生成します

package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

type factory struct{}

func (*factory) Create(ctx context.Context) (Client, error) {
	// S3のエンドポイントを指定するための関数
	endpointFunc := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: env.GetString("AWS_S3_ENDPOINT_URL"),
			}, nil
		},
	)

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     env.GetString("AWS_S3_ACCESS_KEY_ID"),
				SecretAccessKey: env.GetString("AWS_S3_SECRET_ACCESS_KEY"),
			},
		}),
		config.WithEndpointResolverWithOptions(endpointFunc),
	)
	if err != nil {
		return nil, errs.Wrap("[storage.NewClient]AWS接続用の設定読み込みエラー", err)
	}
	return &s3Client{
		client: s3.NewFromConfig(cfg, func(options *s3.Options) {
			// ローカルではUsePathStyleがONになっていないとS3が使えない
			options.UsePathStyle = true
		}),
		bucket: env.GetString("AWS_S3_DELIVERY_BUCKET"),
	}, nil
}
