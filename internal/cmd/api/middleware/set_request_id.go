package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
)

// SetRequestId はリクエスト単位のユニークIDを発行するためのmiddlewareです
type SetRequestId struct{}

// NewSetRequestId はユニークIDを発行するmiddlewareを返します
func NewSetRequestId() *SetRequestId {
	return &SetRequestId{}
}

// Handler はcontextにユニークIDをセットします
// ALB経由でアクセスがあった場合は "X-Amzn-Trace-Id" がヘッダーに存在するのでそちらを利用
// "X-Amzn-Trace-Id" が存在しなければ空文字列を設定
func (*SetRequestId) Handler(ctx *gin.Context) {
	uniqueId := ctx.GetHeader("X-Amzn-Trace-Id")
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), ctxt.KeyLogUniqueId, uniqueId))
}
