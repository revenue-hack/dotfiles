package feature

import (
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
)

func TestHealth(t *testing.T) {
	serv := helper.Server()
	defer serv.Close()

	req, err := http.NewRequest(http.MethodGet, serv.URL+"/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("StatusCode returns %d, want %d", res.StatusCode, http.StatusOK)
	}
}
