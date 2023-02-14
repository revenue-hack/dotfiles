package helper

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

// GetBase64Image はtestdataにある画像をBase64 Encodeした文字列を返却します
func GetBase64Image(t *testing.T, fileName string) string {
	path := fmt.Sprintf("%s/testdata/%s", os.Getenv("TEST_BASE_DIR"), fileName)
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
