package content

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSetting_UrlUpdate(t *testing.T) {
	testCases := []struct {
		name         string
		courseId     int
		contentId    int
		requestBody  string
		statusCode   int
		expectedBody string
		after        func(*testing.T)
	}{
		{
			name:        "存在するURLコンテンツIDを指定した場合、コンテンツ情報が更新される",
			courseId:    1,
			contentId:   31,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusNoContent,
			after: func(t *testing.T) {
				db := helper.OpenDb(t)
				defer helper.CloseDb(t, db)

				assert.EqualFirstRecord(t, db.Where("id = 31"), entity.Content{
					Id:           31,
					CourseId:     1,
					ContentType:  3,
					DisplayOrder: 7,
					CreatedAt:    helper.FixedTime,
					CreatedBy:    1,
					UpdatedAt:    helper.FixedMockTime,
					UpdatedBy:    helper.TestRequestDefaultUserId,
				})
				assert.EqualFirstRecord(t, db.Where("id = 331"), entity.Url{
					Id:        331,
					ContentId: 31,
					Title:     "kaonavi Tech Talk #10",
					Url:       "https://www.youtube.com/watch?v=78t0I4cfDwk",
					CreatedAt: helper.FixedTime,
					CreatedBy: 1,
					UpdatedAt: helper.FixedMockTime,
					UpdatedBy: helper.TestRequestDefaultUserId,
				})
			},
		},

		{
			name:         "タイトルを送信しない場合、422エラーが返却される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["タイトルは必須です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "タイトルに空文字を指定した場合、422エラーが返却される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"title": "", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["タイトルは必須です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:        "タイトルに50文字の文字列を指定した場合、コンテンツが新規作成される",
			courseId:    1,
			contentId:   31,
			requestBody: `{"title": "aAあ🫠漢aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:         "タイトルに51文字の文字列を指定した場合、422エラーが返却される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"title": "aAあ🫠漢aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["タイトルは50文字以内で入力してください"]}`,
			after:        func(t *testing.T) {},
		},

		{
			name:         "URLを送信しない場合、422エラーが返却される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"title": "kaonavi Tech Talk #10"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLは必須です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "URLに空文字を指定した場合、422エラーが返却される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": ""}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLは必須です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:      "URLに255文字の文字列を指定した場合、コンテンツが新規作成される",
			courseId:  1,
			contentId: 31,
			requestBody: fmt.Sprintf(
				`{"title": "kaonavi Tech Talk #10", "url": "%s"}`,
				"https://www.youtube.com/watch?v=78t0I4cfDwk&"+strings.Repeat("a", 211),
			),
			statusCode: http.StatusNoContent,
			after:      func(t *testing.T) {},
		},
		{
			name:      "URLに256文字の文字列を指定した場合、422エラーが返却される",
			courseId:  1,
			contentId: 31,
			requestBody: fmt.Sprintf(
				`{"title": "kaonavi Tech Talk #10", "url": "%s"}`,
				"https://www.youtube.com/watch?v=78t0I4cfDwk&"+strings.Repeat("a", 212),
			),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLは255文字以内で入力してください"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:        "URLのスキームにhttpを指定した場合、コンテンツが新規作成される",
			courseId:    1,
			contentId:   31,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "http://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:         "URLにhttp/https以外のスキームを指定した場合、422エラーが返却される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "file://path/to/file.txt"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLの形式が不正です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "URLにURL形式ではない文字列を指定した場合、422エラーが返却される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "apple pen"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLの形式が不正です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:        "URLのドメインがwww.youtube.comの場合、コンテンツが新規作成される",
			courseId:    1,
			contentId:   31,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "URLのドメインがm.youtube.comの場合、コンテンツが新規作成される",
			courseId:    1,
			contentId:   31,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "https://m.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "URLのドメインがyoutu.beの場合、コンテンツが新規作成される",
			courseId:    1,
			contentId:   31,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "https://youtu.be/78t0I4cfDwk"}`,
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:         "URLに許可されていないドメインが指定されている場合、422エラーが返却される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "https://service.kaonavi.jp"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLに許可されていないドメインが指定されています"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "URLのドメインが有効だがポート番号が付与されている場合、コンテンツが新規作成される",
			courseId:     1,
			contentId:    31,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com:8080/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLに許可されていないドメインが指定されています"]}`,
			after:        func(t *testing.T) {},
		},

		{
			name:         "存在しない講習IDを指定した場合、404エラーが返却される",
			courseId:     999,
			contentId:    31,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "存在しないコンテンツIDを指定した場合、404エラーが返却される",
			courseId:     1,
			contentId:    999,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "URLではないコンテンツIDを指定した場合、404エラーが返却される",
			courseId:     1,
			contentId:    11,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["URLコンテンツが存在しません"]}`,
			after:        func(t *testing.T) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			helper.InitDb(tt)
			helper.ExecSeeder(tt, "setting/content/content")

			res := helper.DoRequest(tt, helper.ApiRequest{
				Method: http.MethodPatch,
				Path:   fmt.Sprintf("/settings/%d/contents/urls/%d", tc.courseId, tc.contentId),
				Body:   tc.requestBody,
			})
			assert.Equal(tt, res.StatusCode, tc.statusCode)
			assert.EqualJson(tt, string(res.Body), tc.expectedBody)
			tc.after(tt)
		})
	}
}
