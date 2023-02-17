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
	t.Run("存在するURLコンテンツIDを指定した場合、コンテンツ情報が削除される", func(tt *testing.T) {
		helper.InitDb(tt, "../testdata/testdata.sql")

		res := helper.DoRequest(tt, helper.ApiRequest{
			Method: http.MethodDelete,
			Path:   fmt.Sprintf("/settings/%d/contents/urls/%d", 1, 31),
		})
		assert.Equal(tt, res.StatusCode, http.StatusNoContent)
		assert.EqualJson(tt, string(res.Body), "")

		db := helper.OpenDb(t)
		defer helper.CloseDb(t, db)

		assert.RecordNotFound(t, db.Where("id = 31"), entity.Content{})
		assert.RecordNotFound(t, db.Where("id = 331"), entity.Url{})
	})

	testCases := []struct {
		name         string
		courseId     int
		contentId    int
		statusCode   int
		expectedBody string
	}{
		{
			name:         "存在しない講習IDを指定した場合、404エラーが返却される",
			courseId:     999,
			contentId:    31,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
		},
		{
			name:         "存在しないコンテンツIDを指定した場合、404エラーが返却される",
			courseId:     1,
			contentId:    999,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
		},
		{
			name:         "URLではないコンテンツIDを指定した場合、404エラーが返却される",
			courseId:     1,
			contentId:    11,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
		},
	}

	t.Run("正常系DB確認無し／エラー系", func(tt *testing.T) {
		helper.InitDb(tt, "../testdata/testdata.sql")

		for _, tc := range testCases {
			tt.Run(tc.name, func(ttt *testing.T) {
				res := helper.DoRequest(ttt, helper.ApiRequest{
					Method: http.MethodDelete,
					Path:   fmt.Sprintf("/settings/%d/contents/urls/%d", tc.courseId, tc.contentId),
				})
				assert.Equal(ttt, res.StatusCode, tc.statusCode)
				if http.StatusNoContent != tc.statusCode {
					// エラー時だけレスポンスを検証します
					assert.EqualJson(ttt, string(res.Body), tc.expectedBody)
				}
			})
		}
	})
}
