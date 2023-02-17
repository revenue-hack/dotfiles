package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

func NewUrlUpdate(uc content.UrlUpdateUseCase) handler.API {
	return &urlUpdate{uc: uc}
}

type urlUpdate struct {
	uc content.UrlUpdateUseCase
}

func (h *urlUpdate) Handler(ctx *gin.Context) {
	handler.Exec(ctx, h.exec)
}

func (h *urlUpdate) exec(ctx *gin.Context) handler.ResponseData {
	courseId, err := vo.NewCourseIdFromString(ctx.Param("course_id"))
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	contentId, err := vo.NewContentIdFromString(ctx.Param("content_id"))
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	var in content.UrlInput
	if err = ctx.ShouldBindJSON(&in); err != nil {
		return handler.NewErrorResponseData(ctx, handler.NewJsonBindError(err))
	}

	if err = h.uc.Exec(ctx, *courseId, *contentId, in); err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	return handler.ResponseData{Status: http.StatusNoContent}
}
