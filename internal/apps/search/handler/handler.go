package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
)

func New(uc search.UseCase) handler.API {
	return &searchHandler{uc: uc}
}

type searchHandler struct {
	uc search.UseCase
}

func (h *searchHandler) Handler(ctx *gin.Context) {
	handler.Exec(ctx, h.exec)
}

func (h *searchHandler) exec(ctx *gin.Context) handler.ResponseData {
	var in search.Input
	if err := ctx.ShouldBindJSON(&in); err != nil {
		return handler.NewErrorResponseData(ctx, handler.NewJsonBindError(err))
	}

	out, err := h.uc.Exec(ctx, in)
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	body, err := h.makeResponse(out)
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	return handler.ResponseData{Status: http.StatusOK, Body: body}
}

func (h *searchHandler) makeResponse(out *search.Output) ([]byte, error) {
	type respThumbnail struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	type respCourse struct {
		Id           uint32         `json:"id"`
		Title        string         `json:"title"`
		Thumbnail    *respThumbnail `json:"thumbnail"`
		CategoryName *string        `json:"categoryName"`
		ExpireAt     *string        `json:"expireAt"`
		Recommend    *uint32        `json:"recommend"`
		IsRequired   bool           `json:"isRequired"`
		IsFixed      bool           `json:"isFixed"`
	}

	resp := struct {
		MaxPageSize   uint32       `json:"maxPageSize"`
		NextPageToken *string      `json:"nextPageToken"`
		Courses       []respCourse `json:"courses"`
	}{
		MaxPageSize:   out.MaxPageSize,
		NextPageToken: out.NextPageToken,
		Courses:       make([]respCourse, 0, len(out.Courses)),
	}

	for _, c := range out.Courses {
		var expireAt *string = nil
		if c.ExpireAt != nil {
			exp := c.ExpireAt.Format("2006/01/02 15:04:05")
			expireAt = &exp
		}
		rc := respCourse{
			Id:           c.Id,
			Title:        c.Title,
			CategoryName: c.CategoryName,
			ExpireAt:     expireAt,
			Recommend:    c.Recommend,
			IsRequired:   c.IsRequired,
			IsFixed:      c.IsFixed,
		}

		if c.Thumbnail != nil {
			rc.Thumbnail = &respThumbnail{
				Name: c.Thumbnail.Name,
				Url:  c.Thumbnail.Url,
			}
		}
		resp.Courses = append(resp.Courses, rc)
	}
	return json.Marshal(&resp)
}
