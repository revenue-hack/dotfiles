package file

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	_ "image/gif"  // image.Decodeに必要
	_ "image/jpeg" // image.Decodeに必要
	_ "image/png"  // image.Decodeに必要
	"net/http"
	"path/filepath"

	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

const (
	imageMinWidth  = 200
	imageMaxWidth  = 2000
	imageMinHeight = 200
	imageMaxHeight = 2000
)

var (
	validImageExt = map[string]struct{}{
		"png": {}, "jpg": {}, "jpeg": {}, "gif": {},
	}
	validImageMimeTypes = map[string]struct{}{
		"image/jpeg": {}, "image/png": {}, "image/gif": {},
	}
)

// NewUploadImage は画像ファイルが正しい場合にUploadFileを返却します
func NewUploadImage(ctx context.Context, name, content string) (*UploadFile, error) {
	f := &UploadFile{
		// 念の為ファイル名だけに変換する
		originalName: filepath.Base(name),
	}

	if _, ok := validImageExt[f.ext()]; !ok {
		return nil, errs.NewInvalidParameter("許可されていない画像の拡張子が指定されています")
	}

	decoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errs.NewInvalidParameter("不正な画像形式のファイルです")
	} else if _, ok := validImageMimeTypes[http.DetectContentType(decoded)]; !ok {
		return nil, errs.NewInvalidParameter("不正な画像形式のファイルです")
	}

	r := bytes.NewReader(decoded)
	img, _, err := image.Decode(r)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errs.NewInvalidParameter("不正な画像形式のファイルです")
	}

	// サイズ・縦横のチェック
	if r.Size() > imageSizeLimit {
		return nil, errs.NewInvalidParameter("%dMB以内の画像を指定してください", imageSizeLimit>>20)
	}
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	if width < imageMinWidth || width > imageMaxWidth || height < imageMinHeight || height > imageMaxHeight {
		return nil, errs.NewInvalidParameter("縦横が200px ~ 2000pxの画像を指定してください")
	}

	f.content = decoded
	return f, nil
}
