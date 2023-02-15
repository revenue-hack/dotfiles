package helper

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

// GetBase64Image はtestdataにある画像をBase64 Encodeした文字列を返却します
func GetBase64Image(t *testing.T, fileName string) string {
	path := fmt.Sprintf("%s/testdata/image/%s", os.Getenv("TEST_BASE_DIR"), fileName)
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

// GetTestStorageFilePath はテスト用のファイル配置先に置かれたファイルを読み込みます
func GetTestStorageFilePath(fileName string) string {
	return fmt.Sprintf("%s/.teststorage/%s/%s", os.Getenv("TEST_BASE_DIR"), dbName, fileName)
}

// CleanTestStorage はテスト用のファイル配置先を掃除します
func CleanTestStorage(t *testing.T) {
	path := fmt.Sprintf("%s/.teststorage/%s", os.Getenv("TEST_BASE_DIR"), dbName)
	if err := os.RemoveAll(path); err != nil {
		t.Error(err)
	}
}
