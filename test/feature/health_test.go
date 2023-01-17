package feature

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/internal/cmd/api"
)

func TestHealth(t *testing.T) {
	serv := httptest.NewServer(api.Route())
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
