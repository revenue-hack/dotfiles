package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/cmd/api/di"
	"gitlab.kaonavi.jp/ae/sardine/internal/cmd/api/middleware"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
)

func Route() *gin.Engine {
	e := gin.New()
	e.ContextWithFallback = true
	gin.SetMode(getMode())

	// アクセスログ
	e.Use(middleware.NewAccessLogWriter().Handler)
	// リカバリログ
	e.Use(middleware.NewRecoveryLogWriter().Handler)
	// リクエストIDの保存
	e.Use(middleware.NewSetRequestId().Handler)
	// タイムアウト
	e.Use(middleware.NewTimeoutHandler().Handler)
	// CORS
	e.Use(cors.New(corsConfig()))

	// ヘルスチェック
	e.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	authGroup := e.Group("", di.InitializeAuthenticateToken().Handler)
	{
		// 検索処理
		authGroup.POST("/search/required", di.InitializeSearchRequiredHandler().Handler)
		authGroup.POST("/search/optional", di.InitializeSearchOptionalHandler().Handler)
		authGroup.POST("/search/expired", di.InitializeSearchExpiredHandler().Handler)
		authGroup.POST("/search/completed", di.InitializeSearchCompletedHandler().Handler)

		// 講習の新規作成
		authGroup.POST("/course/e_learning", di.InitializeCreateELearningHandler().Handler)
	}

	return e
}

func getMode() string {
	mode := env.GetString("APP_MODE")
	if mode == "" {
		return gin.ReleaseMode
	}
	return mode
}

func corsConfig() cors.Config {
	config := cors.Config{
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:       12 * time.Hour,
		// FIXME: 一旦全許可にするがリリース時にはホストを指定するように変更する必要がある
		AllowAllOrigins: true,
	}
	return config
}
