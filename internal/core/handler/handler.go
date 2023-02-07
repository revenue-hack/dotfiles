package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

// API はREST APIのhandler用のインタフェースです
type API interface {
	Handler(*gin.Context)
}

type ResponseData struct {
	Status int
	Body   []byte
}

// Exec はタイムアウトを考慮しHandlerの処理を実行します
func Exec(ctx *gin.Context, handler func(ctx *gin.Context) ResponseData) {
	tx := ctx.Request.Context()
	ch := make(chan ResponseData)
	go func() {
		defer func() {
			// goroutine内でpanicが発生すると親プロセスでrecoverできずに落ちるため
			// goroutine内でrecoverを行う必要があります。
			if err := recover(); err != nil {
				logger.Panic(ctx, err)
				ch <- ResponseData{
					Status: http.StatusInternalServerError,
					Body:   []byte(`{"errors":["Internal Server Error"]}`),
				}
			}
		}()
		ch <- handler(ctx)
	}()

	select {
	case <-tx.Done():
		// タイムアウトの場合はここに来るが、Controllerの処理自体は正常終了させる
		return
	case res := <-ch:
		ctx.Data(res.Status, "application/json; charset=utf-8", res.Body)
	}
}

func NewErrorResponseData(ctx context.Context, err error) ResponseData {
	resp := struct {
		Errors []string `json:"errors"`
	}{
		Errors: []string{},
	}

	// TODO: Switchでエラーの種別を判定してステータスを切り替える
	status := http.StatusBadRequest
	resp.Errors = []string{err.Error()}

	body, err2 := json.Marshal(resp)
	if err2 != nil {
		status = http.StatusInternalServerError
		body = []byte(`{"errors":["Internal Server Error"]}`)
		logger.Error(ctx, err)
	}

	return ResponseData{
		Status: status,
		Body:   body,
	}
}

func NewJsonBindError(err error) error {
	switch err.(type) {
	case *json.UnmarshalTypeError:
		return fmt.Errorf(`%sの型が異なっています。`, err.(*json.UnmarshalTypeError).Field)
	}
	return fmt.Errorf("必須パラメータが不足しています")
}
