package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

func NewList(uc content.ListUseCase) handler.API {
	return &list{uc: uc}
}

type list struct {
	uc content.ListUseCase
}

func (h *list) Handler(ctx *gin.Context) {
	handler.Exec(ctx, h.exec)
}

func (h *list) exec(ctx *gin.Context) handler.ResponseData {
	courseId, err := vo.NewCourseIdFromString(ctx.Param("course_id"))
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	out, err := h.uc.Exec(ctx, *courseId)
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	body, err := h.makeResponse(out)
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	return handler.ResponseData{Status: http.StatusOK, Body: body}
}

func (h *list) makeResponse(out *content.ListOutput) ([]byte, error) {
	type respUrl struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}
	type respContent struct {
		Id   uint32   `json:"id"`
		Type uint8    `json:"type"`
		Url  *respUrl `json:"url,omitempty"`
	}
	resp := struct {
		Contents []respContent `json:"contents"`
	}{
		Contents: make([]respContent, 0, len(out.Contents)),
	}

	for _, c := range out.Contents {
		row := respContent{Id: c.Id, Type: c.ContentType}

		if c.IsUrl() {
			row.Url = &respUrl{Title: c.Url.Title, Url: c.Url.Url}
		} else {
			// 未実装の種別は一旦返さないようにする
			continue
		}

		resp.Contents = append(resp.Contents, row)
	}

	return json.Marshal(&resp)
}
