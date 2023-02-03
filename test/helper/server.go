package helper

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/kgm/auth"
	"gitlab.kaonavi.jp/ae/sardine/internal/cmd/api"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

const (
	TestRequestDefaultCustomerCode = "kaodev"
	TestRequestDefaultUserId       = 123
)

// ApiRequest APIのテストで使用するリクエストパラメータを設定するための構造体
type ApiRequest struct {
	Method string
	Path   string
	Body   string
}

// ApiResponse APIのテストで使用するリクエストパラメータを設定するための構造体
type ApiResponse struct {
	StatusCode int
	Body       []byte
}

// TestToken テスト用のアクセストークン
var TestToken string

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	logger.New(os.Stdout)

	// kgm/authを通す必要があるため、kgmから検証用トークンを発行しておく
	tkn, err := auth.GetTestHS256Token(
		auth.Payload{
			CustomerId:   1,
			CustomerCode: TestRequestDefaultCustomerCode,
			LoginUserId:  TestRequestDefaultUserId,
			LoginAddress: "test@kaodev.jp",
			RoleType:     2,
			Client:       "client",
		},
		auth.Key{
			Encrypt: os.Getenv("TOKEN_ENCRYPT_KEY"),
			Signing: os.Getenv("TOKEN_SIGNING_KEY"),
		},
	)
	if err != nil {
		panic(err)
	}
	TestToken = tkn
}

// DoRequest はAPIリクエストを行います
func DoRequest(t *testing.T, request ApiRequest) ApiResponse {
	t.Helper()

	serv := httptest.NewServer(api.Route())
	defer serv.Close()

	req, err := http.NewRequest(request.Method, serv.URL+request.Path, bytes.NewBuffer([]byte(request.Body)))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+TestToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if e := res.Body.Close(); e != nil {
			t.Fatal(e)
		}
	}()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	return ApiResponse{StatusCode: res.StatusCode, Body: resBody}
}
