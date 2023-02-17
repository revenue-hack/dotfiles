package test

import (
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
)

func TestSearchCompleted(t *testing.T) {
	res := helper.DoRequest(t, helper.ApiRequest{
		Method: http.MethodPost,
		Path:   "/search/completed",
		Body:   `{}`,
	})

	if res.StatusCode != http.StatusOK {
		t.Errorf("StatusCode returns %d, want %d", res.StatusCode, http.StatusOK)
	}

	expected := `{"maxPageSize":0,"nextPageToken":null,"courses":[]}`
	if a := string(res.Body); a != expected {
		t.Errorf("Body returns %s, want %s", a, expected)
	}
}
