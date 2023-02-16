package content

import (
	"fmt"
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSetting_ListContent(t *testing.T) {
	helper.InitDb(t)
	helper.ExecSeeder(t, "setting/content/list")

	testCases := []struct {
		name         string
		courseId     int
		statusCode   int
		expectedBody string
	}{
		{
			name:       "コンテンツが設定されている講習IDを指定した場合、一覧が取得できる",
			courseId:   1,
			statusCode: http.StatusOK,
			expectedBody: `
{
	"contents": [
		{
			"id": 31,
			"type": 3,
			"url": {"title": "kaonavi Tech Talk #1", "url": "https://www.youtube.com/watch?v=HIC4AynFDdw"}
		},
		{
			"id": 33,
			"type": 3,
			"url": {"title": "kaonavi Tech Talk #11", "url": "https://www.youtube.com/watch?v=HPgx7r_I-Ko"}
		},
		{
			"id": 32,
			"type": 3,
			"url": {"title": "kaonavi Tech Talk #2", "url": "https://www.youtube.com/watch?v=3Cs-PVZXsyU"}
		}
	]
}`,
		},
		{
			name:         "コンテンツが設定されていない講習IDを指定した場合、一覧が空で取得できる",
			courseId:     2,
			statusCode:   http.StatusOK,
			expectedBody: `{"contents": []}`,
		},
		{
			name:         "存在しない講習IDを指定した場合、404エラーが返却される",
			courseId:     999,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["講習が存在しません"]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			res := helper.DoRequest(t, helper.ApiRequest{
				Method: http.MethodGet,
				Path:   fmt.Sprintf("/settings/%d/contents", tc.courseId),
			})
			assert.Equal(t, res.StatusCode, tc.statusCode)
			assert.EqualJson(t, string(res.Body), tc.expectedBody)
		})
	}
}
