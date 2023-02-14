package file

import (
	"encoding/base64"
	"net/http"
	"path/filepath"

	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

var (
	validImageExt = map[string]struct{}{
		"png": {}, "jpg": {}, "jpeg": {}, "gif": {},
	}
	validImageMimeTypes = map[string]struct{}{
		"image/jpeg": {}, "image/png": {}, "image/gif": {},
	}
)

// NewUploadImage はアップロードされたファイル情報を扱う構造体を返却します
func NewUploadImage(name, content string) (*UploadFile, error) {
	f := &UploadFile{
		// 念の為ファイル名だけに変換する
		originalName: filepath.Base(name),
	}

	if _, ok := validImageExt[f.ext()]; !ok {
		return nil, errs.NewInvalidParameter("許可されていない画像の拡張子が指定されています")
	}

	decoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, errs.NewInvalidParameter("不正な画像形式のファイルです")
	} else if _, ok := validImageMimeTypes[http.DetectContentType(decoded)]; !ok {
		return nil, errs.NewInvalidParameter("不正な画像形式のファイルです")
	}

	// TODO: ファイルサイズのチェック

	f.content = decoded
	return f, nil
}
