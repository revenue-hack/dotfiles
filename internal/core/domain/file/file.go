package file

import (
	"fmt"
	"path/filepath"
	"strings"

	"gitlab.kaonavi.jp/ae/sardine/internal/utils/hash"
)

type UploadFile struct {
	originalName string
	hashedName   string
	content      []byte
}

func (f *UploadFile) OriginalName() string {
	return f.originalName
}

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

func (f *UploadFile) ext() string {
	return strings.TrimPrefix(strings.ToLower(filepath.Ext(f.originalName)), ".")
}
