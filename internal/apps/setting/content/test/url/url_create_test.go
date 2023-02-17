package url_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSetting_UrlCreate(t *testing.T) {
	t.Run("コンテンツが設定済みの講習IDを指定した場合、コンテンツの並び順が最大数+1で新規作成される", func(tt *testing.T) {
		helper.InitDb(tt, "../testdata/testdata.sql")

		res := helper.DoRequest(tt, helper.ApiRequest{
			Method: http.MethodPost,
			Path:   fmt.Sprintf("/settings/%d/contents/urls", 1),
			Body:   `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
		})
		assert.Equal(tt, res.StatusCode, http.StatusCreated)
		assert.EqualJson(tt, string(res.Body), `{"id": 34}`)

		db := helper.OpenDb(t)
		defer helper.CloseDb(t, db)

		assert.EqualFirstRecord(t, db.Where("id = 34"), entity.Content{
			Id:           34,
			CourseId:     1,
			ContentType:  3,
			DisplayOrder: 10,
			CreatedAt:    helper.FixedMockTime,
			CreatedBy:    helper.TestRequestDefaultUserId,
			UpdatedAt:    helper.FixedMockTime,
			UpdatedBy:    helper.TestRequestDefaultUserId,
		})
		assert.EqualFirstRecord(t, db.Where("id = 334"), entity.Url{
			Id:        334,
			ContentId: 34,
			Title:     "kaonavi Tech Talk #10",
			Url:       "https://www.youtube.com/watch?v=78t0I4cfDwk",
			CreatedAt: helper.FixedMockTime,
			CreatedBy: helper.TestRequestDefaultUserId,
			UpdatedAt: helper.FixedMockTime,
			UpdatedBy: helper.TestRequestDefaultUserId,
		})
	})

	t.Run("コンテンツが設定済みの講習IDを指定した場合、コンテンツの並び順が1で新規作成される", func(tt *testing.T) {
		helper.InitDb(tt, "../testdata/testdata.sql")

		res := helper.DoRequest(tt, helper.ApiRequest{
			Method: http.MethodPost,
			Path:   fmt.Sprintf("/settings/%d/contents/urls", 2),
			Body:   `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
		})
		assert.Equal(tt, res.StatusCode, http.StatusCreated)
		assert.EqualJson(tt, string(res.Body), `{"id": 34}`)

		db := helper.OpenDb(t)
		defer helper.CloseDb(t, db)

		assert.EqualFirstRecord(t, db.Where("id = 34"), entity.Content{
			Id:           34,
			CourseId:     2,
			ContentType:  3,
			DisplayOrder: 1,
			CreatedAt:    helper.FixedMockTime,
			CreatedBy:    helper.TestRequestDefaultUserId,
			UpdatedAt:    helper.FixedMockTime,
			UpdatedBy:    helper.TestRequestDefaultUserId,
		})
		assert.EqualFirstRecord(t, db.Where("id = 334"), entity.Url{
			Id:        334,
			ContentId: 34,
			Title:     "kaonavi Tech Talk #10",
			Url:       "https://www.youtube.com/watch?v=78t0I4cfDwk",
			CreatedAt: helper.FixedMockTime,
			CreatedBy: helper.TestRequestDefaultUserId,
			UpdatedAt: helper.FixedMockTime,
			UpdatedBy: helper.TestRequestDefaultUserId,
		})
	})

	testCases := []struct {
		name         string
		courseId     int
		requestBody  string
		statusCode   int
		expectedBody string
	}{
		{
			name:         "タイトルを送信しない場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["タイトルは必須です"]}`,
		},
		{
			name:         "タイトルに空文字を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["タイトルは必須です"]}`,
		},
		{
			name:        "タイトルに50文字の文字列を指定した場合、コンテンツが新規作成される",
			courseId:    1,
			requestBody: `{"title": "aAあ🫠漢aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusCreated,
		},
		{
			name:         "タイトルに51文字の文字列を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "aAあ🫠漢aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["タイトルは50文字以内で入力してください"]}`,
		},

		{
			name:         "URLを送信しない場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "kaonavi Tech Talk #10"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLは必須です"]}`,
		},
		{
			name:         "URLに空文字を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": ""}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLは必須です"]}`,
		},
		{
			name:     "URLに255文字の文字列を指定した場合、コンテンツが新規作成される",
			courseId: 1,
			requestBody: fmt.Sprintf(
				`{"title": "kaonavi Tech Talk #10", "url": "%s"}`,
				"https://www.youtube.com/watch?v=78t0I4cfDwk&"+strings.Repeat("a", 211),
			),
			statusCode: http.StatusCreated,
		},
		{
			name:     "URLに256文字の文字列を指定した場合、422エラーが返却される",
			courseId: 1,
			requestBody: fmt.Sprintf(
				`{"title": "kaonavi Tech Talk #10", "url": "%s"}`,
				"https://www.youtube.com/watch?v=78t0I4cfDwk&"+strings.Repeat("a", 212),
			),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLは255文字以内で入力してください"]}`,
		},
		{
			name:        "URLのスキームにhttpを指定した場合、コンテンツが新規作成される",
			courseId:    1,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "http://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusCreated,
		},
		{
			name:         "URLにhttp/https以外のスキームを指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "file://path/to/file.txt"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLの形式が不正です"]}`,
		},
		{
			name:         "URLにURL形式ではない文字列を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "apple pen"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLの形式が不正です"]}`,
		},
		{
			name:        "URLのドメインがwww.youtube.comの場合、コンテンツが新規作成される",
			courseId:    1,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusCreated,
		},
		{
			name:        "URLのドメインがm.youtube.comの場合、コンテンツが新規作成される",
			courseId:    1,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "https://m.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:  http.StatusCreated,
		},
		{
			name:        "URLのドメインがyoutu.beの場合、コンテンツが新規作成される",
			courseId:    1,
			requestBody: `{"title": "kaonavi Tech Talk #10", "url": "https://youtu.be/78t0I4cfDwk"}`,
			statusCode:  http.StatusCreated,
		},
		{
			name:         "URLに許可されていないドメインが指定されている場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "https://service.kaonavi.jp"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLに許可されていないドメインが指定されています"]}`,
		},
		{
			name:         "URLのドメインが有効だがポート番号が付与されている場合、コンテンツが新規作成される",
			courseId:     1,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com:8080/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["URLに許可されていないドメインが指定されています"]}`,
		},

		{
			name:         "存在しない講習IDを指定した場合、404エラーが返却される",
			courseId:     999,
			requestBody:  `{"title": "kaonavi Tech Talk #10", "url": "https://www.youtube.com/watch?v=78t0I4cfDwk"}`,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["講習が存在しません"]}`,
		},
	}

	t.Run("正常系DB確認無し／エラー系", func(tt *testing.T) {
		helper.InitDb(tt, "../testdata/testdata.sql")

		for _, tc := range testCases {
			tt.Run(tc.name, func(ttt *testing.T) {
				res := helper.DoRequest(ttt, helper.ApiRequest{
					Method: http.MethodPost,
					Path:   fmt.Sprintf("/settings/%d/contents/urls", tc.courseId),
					Body:   tc.requestBody,
				})
				assert.Equal(ttt, res.StatusCode, tc.statusCode)
				if http.StatusCreated != tc.statusCode {
					// エラー時だけレスポンスを検証します
					assert.EqualJson(ttt, string(res.Body), tc.expectedBody)
				}
			})
		}
	})
}
