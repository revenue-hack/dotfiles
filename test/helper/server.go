package helper

import (
	"io"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/cmd/api"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	logger.New(os.Stdout)
}

func Server() *httptest.Server {
	return httptest.NewServer(api.Route())
}
