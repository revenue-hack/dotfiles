package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

type TimeoutHandler struct{}

func NewTimeoutHandler() *TimeoutHandler {
	return &TimeoutHandler{}
}

func (*TimeoutHandler) Handler(ctx *gin.Context) {
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
				// TODO: errsパッケージを作ったら置き換える
				fmt.Errorf("タイムアウトが発生しました %s %s", ctx.Request.Method, ctx.Request.URL.String()),
			)

			// ここに到達時点でレスポンスは出力済みのはずなのでAbortするだけにする
			ctx.Abort()
		}
		cancel()
	}()

	ctx.Request = ctx.Request.WithContext(tx)
	ctx.Next()
}
