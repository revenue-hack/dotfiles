package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/timer"
)

// ログに出力するUserAgentの最大文字数
const userAgentMaxLength = 150

// AccessLogWriter はアクセスログの出力を行います
type AccessLogWriter struct {
	skip map[string]struct{}
}

// NewAccessLogWriter はAccessLogWriterを返します
func NewAccessLogWriter() *AccessLogWriter {
	// ログ出力したくないパスの定義
	skipPaths := []string{
		"/favicon.ico",
	}
	skip := make(map[string]struct{}, len(skipPaths))
	for _, path := range skipPaths {
		skip[path] = struct{}{}
	}

	return &AccessLogWriter{
		skip: skip,
	}
}

// Handler はアクセスログを出力します
// https://github.com/gin-gonic/gin/blob/master/logger.go の LoggerWithConfig の処理を一部移植
func (r *AccessLogWriter) Handler(ctx *gin.Context) {
	// Start timer
	start := timer.Now()
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery

	// Process request
	ctx.Next()

	// Log only when path is not being skipped
	if _, ok := r.skip[path]; !ok {
		// Stop timer
		timeStamp := timer.Now()
		latency := timeStamp.Sub(start)

		if raw != "" {
			path = fmt.Sprintf("%s?%s", path, raw)
		}
		if latency > time.Minute {
			latency = latency - latency%time.Second
		}

		logger.Info(
			ctx,
			logger.MakeMessage("access log"),
			logger.MakeParameter("status", ctx.Writer.Status()),
			logger.MakeParameter("latency", latency.Milliseconds()),
			logger.MakeParameter("ip", ctx.ClientIP()),
			logger.MakeParameter("response_size", ctx.Writer.Size()),
			logger.MakeParameter("method", ctx.Request.Method),
			logger.MakeParameter("path", path),
			logger.MakeParameter("user_agent", getUserAgent(ctx)),
		)
	}
}

// マルチバイトを考慮して150文字以上の場合はUserAgentを削る
func getUserAgent(ctx *gin.Context) *string {
	ua, ok := ctx.Request.Header["User-Agent"]
	if !ok {
		return (*string)(nil)
	}

	// マルチバイトの場合も考慮してキャストする
	target := []rune(ua[0])
	if len(target) > userAgentMaxLength {
		target = target[0:userAgentMaxLength]
	}

	result := string(target)
	return &result
}
