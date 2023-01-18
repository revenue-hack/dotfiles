package middleware

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

// RecoveryLogWriter はリカバリログの出力を行います
type RecoveryLogWriter struct{}

// NewRecoveryLogWriter はRecoveryLogWriterを返します
func NewRecoveryLogWriter() *RecoveryLogWriter {
	return &RecoveryLogWriter{}
}

// Handler はリカバリログを出力します
// https://github.com/gin-gonic/gin/blob/master/recovery.go の RecoveryWithWriter の処理を一部移植
func (r *RecoveryLogWriter) Handler(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		httpRequest, _ := httputil.DumpRequest(ctx.Request, false)

		if r.isSyscallError(err) {
			logger.Error(
				ctx,
				// TODO: errsパッケージを作ったら置き換える
				fmt.Errorf("%v", err),
				logger.MakeParameter("request", string(httpRequest)),
			)
			ctx.Error(err.(error)) // nolint: errcheck
			ctx.Abort()
			return
		}

		if gin.IsDebugging() {
			headers := strings.Split(string(httpRequest), "\n")
			for idx, header := range headers {
				current := strings.Split(header, ":")
				if current[0] == "Authorization" {
					headers[idx] = current[0] + ": *"
				}
			}
			logger.Panic(ctx, err, logger.MakeParameter("headers", headers))
		} else {
			logger.Panic(ctx, err)
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"Internal Server Error"}})
		ctx.Abort()
	}()

	ctx.Next()
}

// Check for a broken connection, as it is not really a
// condition that warrants a panic stack trace.
func (r *RecoveryLogWriter) isSyscallError(err any) bool {
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
				strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				return true
			}
		}
	}
	return false
}
