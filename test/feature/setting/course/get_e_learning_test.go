package course

import (
	"fmt"
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSetting_GetELearning(t *testing.T) {
	helper.InitDb(t)
	helper.ExecSeeder(t, "setting/course/course")

	testCases := []struct {
		name         string
		courseId     int
		statusCode   int
		expectedBody string
	}{
		{
			name:       "期間が存在する講習IDを指定した場合",
			courseId:   1,
			statusCode: http.StatusOK,
			expectedBody: fmt.Sprintf(`
{
	"id": 1,
	"title": "e-Learning 1",
	"description": "e-Learningの説明1",
	"thumbnail": {"name": "original1.png", "url": "%s"},
	"isRequired": false,
	"categoryId": 1,
	"from": "2023/02/01 10:00",
	"to": "2023/03/01 18:30"
}`, helper.GetThumbnailDeliveryUrl(1, "delivery1.png")),
		},
		{
			name:       "期間が存在しない講習IDを指定した場合、From/Toはnullが返却される",
			courseId:   2,
			statusCode: http.StatusOK,
			expectedBody: `
{
	"id": 2,
	"title": "e-Learning 2",
	"description": "",
	"thumbnail": null,
	"isRequired": true,
	"categoryId": 2,
	"from": null,
	"to": null
}`,
		},
		{
			name:         "存在しない講習IDを指定した場合、404エラーが返却される",
			courseId:     999,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["講習が存在しません"]}`,
		},
		{
			name:         "講習区分がe-Learning以外の講習IDを指定した場合、404エラーが返却される",
			courseId:     3,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["講習が存在しません"]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			res := helper.DoRequest(tt, helper.ApiRequest{
				Method: http.MethodGet,
				Path:   fmt.Sprintf("/settings/%d/e_learning", tc.courseId),
			})
			assert.Equal(tt, res.StatusCode, tc.statusCode)
			assert.EqualJson(tt, string(res.Body), tc.expectedBody)
		})
	}

}
