package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

func NewUrlCreate(uc content.UrlCreateUseCase) handler.API {
	return &urlCreate{uc: uc}
}

type urlCreate struct {
	uc content.UrlCreateUseCase
}

func (h *urlCreate) Handler(ctx *gin.Context) {
	handler.Exec(ctx, h.exec)
}

func (h *urlCreate) exec(ctx *gin.Context) handler.ResponseData {
	courseId, err := vo.NewCourseIdFromString(ctx.Param("course_id"))
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	var in content.UrlInput
	if err = ctx.ShouldBindJSON(&in); err != nil {
		return handler.NewErrorResponseData(ctx, handler.NewJsonBindError(err))
	}

	out, err := h.uc.Exec(ctx, *courseId, in)
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	body, err := h.makeResponse(out)
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	return handler.ResponseData{Status: http.StatusCreated, Body: body}
}

func (h *urlCreate) makeResponse(out *content.CreateOutput) ([]byte, error) {
	resp := struct {
		Id uint32 `json:"id"`
	}{
		Id: out.ContentId,
	}
	return json.Marshal(&resp)
}
