package course

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
		Path:   "/course/e_learning",
		Body:   `{}`,
	})
	expected := `{"id": 1}`

	assert.Equal(t, res.StatusCode, http.StatusCreated)
	assert.EqualJson(t, string(res.Body), expected)

	// DBに保存されたデータを検証
	db := helper.OpenDb(t)
	defer helper.CloseDb(t, db)

	expectedCourse := entity.Course{
		Id:                 1,
		CourseType:         1,
		Title:              "無題のe-Learning",
		Description:        nil,
		ThumbnailImageName: nil,
		IsRequired:         false,
		CategoryId:         nil,
		Status:             1,
		CreatedAt:          helper.FixedMockTime,
		CreatedBy:          helper.TestRequestDefaultUserId,
		UpdatedAt:          helper.FixedMockTime,
		UpdatedBy:          helper.TestRequestDefaultUserId,
	}

	var actualCourse entity.Course
	if err := db.First(&actualCourse).Error; err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, actualCourse, expectedCourse)
}
