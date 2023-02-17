package course_test

import (
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestCreateELearning(t *testing.T) {
	helper.InitDb(t)

	res := helper.DoRequest(t, helper.ApiRequest{
		Method: http.MethodPost,
		Path:   "/courses/e_learning",
		Body:   `{}`,
	})
	expected := `{"id": 1}`

	assert.Equal(t, res.StatusCode, http.StatusCreated)
	assert.EqualJson(t, string(res.Body), expected)

	// DBに保存されたデータを検証
	db := helper.OpenDb(t)
	defer helper.CloseDb(t, db)

	assert.EqualFirstRecord(t, db, entity.Course{
		Id:                        1,
		CourseType:                1,
		Title:                     "無題のe-Learning",
		Description:               nil,
		ThumbnailDeliveryFileName: nil,
		ThumbnailOriginalFileName: nil,
		IsRequired:                false,
		CategoryId:                nil,
		Status:                    1,
		CreatedAt:                 helper.FixedMockTime,
		CreatedBy:                 helper.TestRequestDefaultUserId,
		UpdatedAt:                 helper.FixedMockTime,
		UpdatedBy:                 helper.TestRequestDefaultUserId,
	})
}
