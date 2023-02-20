package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/attend"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

func NewGetELearning(uc attend.GetELearningUseCase) handler.API {
	return &getELearning{uc: uc}
}

type getELearning struct {
	uc attend.GetELearningUseCase
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

func (h *getELearning) makeResponse(out *attend.GetELearningOutput) ([]byte, error) {
	type respForm struct {
		Id              int    `json:"id"`
		Title           string `json:"title"`
		AnsweredCount   int    `json:"answeredCount"`
		AnswerLimit     int    `json:"answerLimit"`
		AnswerTimeLimit int    `json:"answerTimeLimit"`
		PassPoint       int    `json:"passPoint"`
	}
	type respContentUrl struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}
	type respContent struct {
		Id   uint32          `json:"id"`
		Type uint8           `json:"type"`
		Url  *respContentUrl `json:"url,omitempty"`
	}

	resp := struct {
		Title       string        `json:"title"`
		Description *string       `json:"description"`
		From        *string       `json:"from"`
		To          *string       `json:"to"`
		IsFixed     bool          `json:"isFixed"`
		HasForm     bool          `json:"hasForm"`
		Form        *respForm     `json:"form"`
		Contents    []respContent `json:"contents"`
	}{
		Title:       out.Title,
		Description: out.Description,
		IsFixed:     out.IsFixed,
		Contents:    make([]respContent, 0, len(out.Contents)),
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

	for _, c := range out.Contents {
		row := respContent{Id: c.Id, Type: c.ContentType}
		if c.IsUrl() {
			row.Url = &respContentUrl{Title: c.Url.Title, Url: c.Url.Url}
		} else {
			// 未実装の種別は一旦返さないようにする
			continue
		}
		resp.Contents = append(resp.Contents, row)
	}

	return json.Marshal(&resp)
}
