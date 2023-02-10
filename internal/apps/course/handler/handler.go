package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
)

func NewCreateHandler(uc course.CreateUseCase) handler.API {
	return &createHandler{uc: uc}
}

type createHandler struct {
	uc course.CreateUseCase
}

func (h *createHandler) Handler(ctx *gin.Context) {
	handler.Exec(ctx, h.exec)
}

func (h *createHandler) exec(ctx *gin.Context) handler.ResponseData {
	out, err := h.uc.Exec(ctx)
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	body, err := h.makeResponse(out)
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	return handler.ResponseData{Status: http.StatusCreated, Body: body}
}

func (h *createHandler) makeResponse(out *course.CreateOutput) ([]byte, error) {
	resp := struct {
		Id uint32 `json:"id"`
	}{
		Id: out.Course.Id,
	}
	return json.Marshal(&resp)
}
