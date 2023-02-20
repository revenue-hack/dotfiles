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
	// ヘルスチェック
	e.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// CORS
	e.Use(cors.New(corsConfig()))

	authGroup := e.Group("", di.InitializeAuthenticateToken().Handler)
	{
		// 検索処理
		authGroup.POST("/courses/search/required", di.InitializeSearchRequiredHandler().Handler)
		authGroup.POST("/courses/search/optional", di.InitializeSearchOptionalHandler().Handler)
		authGroup.POST("/courses/search/expired", di.InitializeSearchExpiredHandler().Handler)
		authGroup.POST("/courses/search/completed", di.InitializeSearchCompletedHandler().Handler)

		// 受講画面
		authGroup.GET("/attend/:course_id/e_learning", di.InitializeAttendGetELearningHandler().Handler)

		// TODO: これ以降のルーティングは編集権限を持つ場合のみ利用可能なので、別のgroupにして別途ミドルウェアを設定

		// 講習の新規作成
		authGroup.POST("/courses/e_learning", di.InitializeCreateELearningHandler().Handler)

		// 講習内容の編集
		// 概要取得・更新
		authGroup.GET("/settings/:course_id/e_learning", di.InitializeSettingGetELearningHandler().Handler)
		authGroup.PATCH("/settings/:course_id/e_learning", di.InitializeSettingUpdateELearningHandler().Handler)
		// コンテンツ管理
		authGroup.GET("/settings/:course_id/contents", di.InitializeSettingListContentHandler().Handler)
		authGroup.POST("/settings/:course_id/contents/urls", di.InitializeSettingUrlCreateHandler().Handler)
		authGroup.PATCH("/settings/:course_id/contents/urls/:content_id", di.InitializeSettingUrlUpdateHandler().Handler)
		authGroup.DELETE("/settings/:course_id/contents/urls/:content_id", di.InitializeSettingUrlDeleteHandler().Handler)
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
