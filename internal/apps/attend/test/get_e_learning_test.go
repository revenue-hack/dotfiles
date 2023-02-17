package test

import (
	"fmt"
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestAttend_GetELearning(t *testing.T) {
	helper.InitDb(t, "testdata/testdata.sql")

	testCases := []struct {
		name         string
		courseId     int
		statusCode   int
		expectedBody string
	}{
		{
			name:       "期間、コンテンツが存在する講習IDを指定した場合",
			courseId:   1,
			statusCode: http.StatusOK,
			expectedBody: `
{
	"title": "e-Learning 1",
	"description": "e-Learningの説明1",
	"from": "2023/02/01 10:00",
	"to": "2023/03/01 18:30",
	"isFixed": false,
	"hasForm": false,
	"form": null,
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
			name:       "期間、コンテンツが存在しない講習IDを指定した場合、From/Toはnullが返却される",
			courseId:   2,
			statusCode: http.StatusOK,
			expectedBody: `
{
	"title": "e-Learning 2",
	"description": "",
	"from": null,
	"to": null,
	"isFixed": false,
	"hasForm": false,
	"form": null,
	"contents": []
}`,
		},
		{
			name:         "ステータスが非公開の講習IDを指定した場合、404エラーが返却される",
			courseId:     3,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["講習が存在しません"]}`,
		},
		{
			name:         "講習区分がe-Learning以外の講習IDを指定した場合、404エラーが返却される",
			courseId:     4,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["講習が存在しません"]}`,
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
			res := helper.DoRequest(tt, helper.ApiRequest{
				Method: http.MethodGet,
				Path:   fmt.Sprintf("/attend/%d/e_learning", tc.courseId),
			})
			assert.Equal(tt, res.StatusCode, tc.statusCode)
			assert.EqualJson(tt, string(res.Body), tc.expectedBody)
		})
	}
}
