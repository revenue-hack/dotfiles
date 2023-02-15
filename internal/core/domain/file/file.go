package file

import (
	"fmt"
	"path/filepath"
	"strings"

	"gitlab.kaonavi.jp/ae/sardine/internal/utils/hash"
)

// UploadFile はアップロードされたファイル情報を扱うための構造体です
type UploadFile struct {
	originalName string
	hashedName   string
	content      []byte
}

// OriginalName は元ファイル名を返却します
func (f *UploadFile) OriginalName() string {
	return f.originalName
}

// Content はファイル本体のbyte列を返却します
func (f *UploadFile) Content() []byte {
	return f.content
}

// HashedName ハッシュ化したファイル名を返却します
func (f *UploadFile) HashedName() string {
	if f.hashedName == "" {
		f.hashedName = fmt.Sprintf("%s.%s", hash.Make(f.originalName), f.ext())
	}
	return f.hashedName
}

// ext は拡張子のみを返却します
func (f *UploadFile) ext() string {
	return strings.TrimPrefix(strings.ToLower(filepath.Ext(f.originalName)), ".")
}
