package url

import (
	"fmt"
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSetting_UrlDelete(t *testing.T) {
	testCases := []struct {
		name         string
		courseId     int
		contentId    int
		statusCode   int
		expectedBody string
		after        func(*testing.T)
	}{
		{
			name:       "存在するURLコンテンツIDを指定した場合、コンテンツ情報が削除される",
			courseId:   1,
			contentId:  31,
			statusCode: http.StatusNoContent,
			after: func(t *testing.T) {
				db := helper.OpenDb(t)
				defer helper.CloseDb(t, db)

				assert.RecordNotFound(t, db.Where("id = 31"), entity.Content{})
				assert.RecordNotFound(t, db.Where("id = 331"), entity.Url{})
			},
		},

		{
			name:         "存在しない講習IDを指定した場合、404エラーが返却される",
			courseId:     999,
			contentId:    31,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "存在しないコンテンツIDを指定した場合、404エラーが返却される",
			courseId:     1,
			contentId:    999,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "URLではないコンテンツIDを指定した場合、404エラーが返却される",
			courseId:     1,
			contentId:    11,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
			after:        func(t *testing.T) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			helper.InitDb(tt)
			helper.InitDb(t, "../testdata/testdata.sql")

			res := helper.DoRequest(tt, helper.ApiRequest{
				Method: http.MethodDelete,
				Path:   fmt.Sprintf("/settings/%d/contents/urls/%d", tc.courseId, tc.contentId),
			})
			assert.Equal(tt, res.StatusCode, tc.statusCode)
			assert.EqualJson(tt, string(res.Body), tc.expectedBody)
			tc.after(tt)
		})
	}
}
