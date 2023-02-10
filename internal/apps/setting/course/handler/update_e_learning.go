package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/vo"
)

func NewUpdateELearning(uc course.UpdateELearningUseCase) handler.API {
	return &updateELearning{uc: uc}
}

type updateELearning struct {
	uc course.UpdateELearningUseCase
}

func (h *updateELearning) Handler(ctx *gin.Context) {
	handler.Exec(ctx, h.exec)
}

func (h *updateELearning) exec(ctx *gin.Context) handler.ResponseData {
	courseId, err := vo.NewCourseIdFromString(ctx.Param("course_id"))
	if err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}

	var in course.UpdateELearningInput
	if err = ctx.ShouldBindJSON(&in); err != nil {
		return handler.NewErrorResponseData(ctx, handler.NewJsonBindError(err))
	}

	if err = h.uc.Exec(ctx, *courseId, in); err != nil {
		return handler.NewErrorResponseData(ctx, err)
	}
	return handler.ResponseData{Status: http.StatusNoContent, Body: nil}
}
