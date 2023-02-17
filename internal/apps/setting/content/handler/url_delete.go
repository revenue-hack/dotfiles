package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

func NewUrlDelete(uc content.UrlDeleteUseCase) handler.API {
	return &urlDelete{uc: uc}
}

type urlDelete struct {
	uc content.UrlDeleteUseCase
}

func (h *urlDelete) Handler(ctx *gin.Context) {
	handler.Exec(ctx, h.exec)
}

func (h *urlDelete) exec(ctx *gin.Context) handler.ResponseData {
	courseId, err := vo.NewCourseIdFromString(ctx.Param("course_id"))
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	contentId, err := vo.NewContentIdFromString(ctx.Param("content_id"))
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	if err = h.uc.Exec(ctx, *courseId, *contentId); err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	return handler.ResponseData{Status: http.StatusNoContent}
}
