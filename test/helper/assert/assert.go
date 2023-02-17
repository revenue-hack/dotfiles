package assert

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"
)

// Equal は値が等しいことを検証します
func Equal(t *testing.T, actual, expected any) {
	t.Helper()
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("values must match:\n%s", diff)
	}
}

// NotEqual は値が等しくないことを検証します
func NotEqual(t *testing.T, actual, expected any) {
	t.Helper()
	if ok := cmp.Equal(actual, expected); ok {
		t.Errorf("values must not match: %v", expected)
	}
}

// EqualJson はJSON値同士が等しいことを検証します
func EqualJson(t *testing.T, actual, expected string) {
	t.Helper()
	compactJson := func(s string) string {
		if s == "" {
			return s
		}

		buf := bytes.NewBuffer(nil)
		if err := json.Compact(buf, []byte(s)); err != nil {
			t.Error(err)
			return s
		} else {
			return buf.String()
		}
	}

	if diff := cmp.Diff(compactJson(actual), compactJson(expected)); diff != "" {
		t.Errorf("json values must match:\n%s", diff)
	}
}

// EqualFirstRecord は指定条件で検索した結果の1レコード目と期待値が等しいことを検証します
func EqualFirstRecord[T any](t *testing.T, query *gorm.DB, expected T) {
	t.Helper()
	var actual T
	if err := query.First(&actual).Error; err != nil {
		t.Error(err)
		return
	}
	Equal(t, actual, expected)
}

// RecordNotFound は指定条件で検索した結果が存在しないことを検証します
func RecordNotFound[T any](t *testing.T, query *gorm.DB, model T) {
	t.Helper()
	if err := query.First(&model).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			t.Error(err)
			return
		}
		// レコード無しは正常なので終了
		return
	}
	t.Error("record should not exist")
}

// FileExist はファイルが存在することを検証します
func FileExist(t *testing.T, path string) {
	sts, err := os.Stat(path)
	if err != nil {
		t.Error(err)
		return
	}
	if sts.IsDir() {
		t.Errorf("path %q is a directory", path)
	}
}
