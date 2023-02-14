package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

func NewGetELearning(uc course.GetELearningUseCase) handler.API {
	return &getELearning{uc: uc}
}

type getELearning struct {
	uc course.GetELearningUseCase
}

func (h *getELearning) Handler(ctx *gin.Context) {
	handler.Exec(ctx, h.exec)
}

func (h *getELearning) exec(ctx *gin.Context) handler.ResponseData {
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

func (h *getELearning) makeResponse(out *course.GetELearningOutput) ([]byte, error) {
	type thumbnail struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	resp := struct {
		Id          uint32     `json:"id"`
		Title       string     `json:"title"`
		Description *string    `json:"description"`
		Thumbnail   *thumbnail `json:"thumbnail"`
		IsRequired  bool       `json:"isRequired"`
		CategoryId  *uint32    `json:"categoryId"`
		From        *string    `json:"from"`
		To          *string    `json:"to"`
	}{
		Id:          out.Id,
		Title:       out.Title,
		Description: out.Description,
		IsRequired:  out.IsRequired,
		CategoryId:  out.CategoryId,
	}

	format := func(t time.Time) *string {
		s := t.Format("2006/01/02 15:04")
		return &s
	}
	if out.From != nil {
		resp.From = format(*out.From)
	}
	if out.To != nil {
		resp.To = format(*out.To)
	}

	if out.Thumbnail != nil {
		resp.Thumbnail = &thumbnail{
			Name: out.Thumbnail.Name,
			Url:  out.Thumbnail.Url,
		}
	}

	return json.Marshal(&resp)
}
