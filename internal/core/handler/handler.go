package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

const internalError = `{"errors":["internal server error"]}`

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
				ch <- ResponseData{Status: http.StatusInternalServerError, Body: []byte(internalError)}
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
	var status int
	resp := struct {
		Errors []string `json:"errors"`
	}{
		Errors: []string{},
	}

	var ers *errs.Errors
	switch {
	case errors.As(err, &ers):
		status = http.StatusUnprocessableEntity
		resp.Errors = ers.Errors()
	case errors.Is(err, &errs.InvalidParameter{}):
		status = http.StatusUnprocessableEntity
		resp.Errors = append(resp.Errors, err.Error())
	case errors.Is(err, &errs.NotFound{}):
		status = http.StatusNotFound
		resp.Errors = append(resp.Errors, err.Error())
	default:
		status = http.StatusInternalServerError
		resp.Errors = append(resp.Errors, "internal server error")
	}

	logger.Error(ctx, err)
	body, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		// jsonエンコードが失敗することはないはずだが一応失敗したら固定文言を返すようにしておく
		status = http.StatusInternalServerError
		body = []byte(internalError)
		logger.Error(ctx, marshalErr)
	}
	return ResponseData{Status: status, Body: body}
}

func NewJsonBindError(err error) error {
	switch err.(type) {
	case *json.UnmarshalTypeError:
		return errs.NewInvalidParameter(`%sの型が異なっています。`, err.(*json.UnmarshalTypeError).Field)
	}
	return errs.NewInvalidParameter("必須パラメータが不足しています")
}
