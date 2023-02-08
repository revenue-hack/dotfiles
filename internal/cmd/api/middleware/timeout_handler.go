package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

type TimeoutHandler struct{}

func NewTimeoutHandler() *TimeoutHandler {
	return &TimeoutHandler{}
}

func (*TimeoutHandler) Handler(ctx *gin.Context) {
	// TODO: タイムアウト時間が長すぎるかもしれないので、後で数値を調整する
	// 5m30s: ALBのタイムアウト値(300) < アプリのタイムアウト値にするため
	timeout := time.Second * 330
	tx, cancel := context.WithTimeout(ctx.Request.Context(), timeout)
	defer func() {
		if ctx.IsAborted() {
			return
		}
		if tx.Err() == context.DeadlineExceeded {
			logger.Error(
				ctx,
				errs.NewInternalError("タイムアウトが発生しました %s %s", ctx.Request.Method, ctx.Request.URL.String()),
			)

			// ここに到達時点でレスポンスは出力済みのはずなのでAbortするだけにする
			ctx.Abort()
		}
		cancel()
	}()

	ctx.Request = ctx.Request.WithContext(tx)
	ctx.Next()
}
