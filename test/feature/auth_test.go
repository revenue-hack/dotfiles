// TODO: 一時的な検証用のファイルなので後で消します

package feature

import (
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
)

func TestAuth(t *testing.T) {
	res := helper.DoRequest(t, helper.ApiRequest{
		Method: http.MethodGet,
		Path:   "/test",
	})

	if res.StatusCode != http.StatusOK {
		t.Errorf("StatusCode returns %d, want %d", res.StatusCode, http.StatusOK)
	}

	expected := `{"message":"ok"}`
	if a := string(res.Body); a != expected {
		t.Errorf("Body returns %s, want %s", a, expected)
	}
}
