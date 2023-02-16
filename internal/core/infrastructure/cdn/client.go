//go:build !local && !feature_test

package cdn

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/timer"
)

func NewFactory() Factory {
	return &factory{}
}

type factory struct{}

func (*factory) Create(ctx context.Context) (Client, error) {
	return &client{
		baseUrl: env.GetString("AWS_CDN_BASE_URL"),
		// TODO: demo環境ができる頃に実装する
		// signer:  sign.NewURLSigner(keyID, privKey),
	}, nil
}

type client struct {
	baseUrl string
	signer  *sign.URLSigner
}

func (c *client) CreateUrl(path string) (string, error) {
	rawUrl := fmt.Sprintf("%s/%s", c.baseUrl, strings.TrimPrefix(path, "/"))
	url, err := c.signer.Sign(rawUrl, timer.Now().Add(1*time.Hour))
	if err != nil {
		return "", errs.Wrap("[cdn.client.CreateUrl]署名付きURLの作成に失敗", err)
	}
	return url, nil
}
