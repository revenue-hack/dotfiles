//go:build !feature_test

package storage

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

type s3Client struct {
	client *s3.Client
	bucket string
}

func (c *s3Client) Create(ctx context.Context, distPath string, content []byte) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(distPath),
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		return errs.Wrap("[s3Client.Create]PutObjectのエラー", err)
	}
	return nil
}
