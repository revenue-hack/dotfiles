package course

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/internal/entity"
	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSetting_UpdateELearning(t *testing.T) {
	helper.InitDb(t)
	helper.ExecSeeder(t, "setting/course/course")

	testCases := []struct {
		name         string
		courseId     int
		requestBody  string
		statusCode   int
		expectedBody string
		after        func(*testing.T)
	}{
		{
			name:     "スケジュールが登録済みの講習に対して有効な値を送信した場合、講習情報が更新できる",
			courseId: 1,
			requestBody: `
{
	"title": "ロジカルシンキング研修",
	"description": "ロジカルな思考をゲット",
	"thumbnail": null,
	"isRemoveThumbnailImage": false,
	"isRequired": true,
	"categoryId": 2,
	"from": "2024/02/01 10:00",
	"to": "2024/03/01 18:30"
}`,
			statusCode: http.StatusNoContent,
			after: func(t *testing.T) {
				// DBに保存されたデータを検証
				db := helper.OpenDb(t)
				defer helper.CloseDb(t, db)

				// 講習の値が等しいこと
				assert.EqualFirstRecord(t, db.Where("id = 1"), entity.Course{
					Id:                        1,
					CourseType:                1,
					Title:                     "ロジカルシンキング研修",
					Description:               helper.P("ロジカルな思考をゲット"),
					ThumbnailDeliveryFileName: helper.P("delivery1.png"),
					ThumbnailOriginalFileName: helper.P("original1.png"),
					IsRequired:                true,
					CategoryId:                helper.P[uint32](2),
					Status:                    1,
					CreatedAt:                 helper.FixedTime,
					CreatedBy:                 1,
					UpdatedAt:                 helper.FixedMockTime,
					UpdatedBy:                 helper.TestRequestDefaultUserId,
				})
				// 講習に紐づくスケジュール（親）が等しいこと
				assert.EqualFirstRecord(t, db.Where("course_id = 1"), entity.CourseSchedule{
					Id:        10,
					CourseId:  1,
					CreatedAt: helper.FixedTime,
					CreatedBy: 1,
					UpdatedAt: helper.FixedMockTime,
					UpdatedBy: helper.TestRequestDefaultUserId,
				})
				// 講習に紐づくe-Learningのスケジュール（子）が等しいこと
				assert.EqualFirstRecord(t, db.Where("course_schedule_id = 10"), entity.ELearningSchedule{
					Id:               100,
					CourseScheduleId: 10,
					From:             helper.P(helper.AToTime(t, "2024-02-01 10:00:00")),
					To:               helper.P(helper.AToTime(t, "2024-03-01 18:30:00")),
					CreatedAt:        helper.FixedTime,
					CreatedBy:        1,
					UpdatedAt:        helper.FixedMockTime,
					UpdatedBy:        helper.TestRequestDefaultUserId,
				})
			},
		},
		{
			name:     "スケジュールが登録されていない講習に対して有効な値を送信した場合、講習情報が更新できる",
			courseId: 2,
			requestBody: `
{
	"title": "ロジカルシンキング研修",
	"description": "ロジカルな思考をゲット",
	"thumbnail": null,
	"isRemoveThumbnailImage": false,
	"isRequired": false,
	"categoryId": 1,
	"from": "2024/02/01 10:00",
	"to": "2024/03/01 18:30"
}`,
			statusCode: http.StatusNoContent,
			after: func(t *testing.T) {
				// DBに保存されたデータを検証
				db := helper.OpenDb(t)
				defer helper.CloseDb(t, db)

				// 講習の値が等しいこと
				assert.EqualFirstRecord(t, db.Where("id = 2"), entity.Course{
					Id:                        2,
					CourseType:                1,
					Title:                     "ロジカルシンキング研修",
					Description:               helper.P("ロジカルな思考をゲット"),
					ThumbnailDeliveryFileName: nil,
					ThumbnailOriginalFileName: nil,
					IsRequired:                false,
					CategoryId:                helper.P[uint32](1),
					Status:                    2,
					CreatedAt:                 helper.FixedTime,
					CreatedBy:                 1,
					UpdatedAt:                 helper.FixedMockTime,
					UpdatedBy:                 helper.TestRequestDefaultUserId,
				})
				// 講習に紐づくスケジュール（親）が新規追加されていて等しいこと
				assert.EqualFirstRecord(t, db.Where("course_id = 2"), entity.CourseSchedule{
					Id:        31,
					CourseId:  2,
					CreatedAt: helper.FixedMockTime,
					CreatedBy: helper.TestRequestDefaultUserId,
					UpdatedAt: helper.FixedMockTime,
					UpdatedBy: helper.TestRequestDefaultUserId,
				})
				// 講習に紐づくe-Learningのスケジュール（子）が新規追加されていて等しいこと
				assert.EqualFirstRecord(t, db.Where("course_schedule_id = 31"), entity.ELearningSchedule{
					Id:               301,
					CourseScheduleId: 31,
					From:             helper.P(helper.AToTime(t, "2024-02-01 10:00:00")),
					To:               helper.P(helper.AToTime(t, "2024-03-01 18:30:00")),
					CreatedAt:        helper.FixedMockTime,
					CreatedBy:        helper.TestRequestDefaultUserId,
					UpdatedAt:        helper.FixedMockTime,
					UpdatedBy:        helper.TestRequestDefaultUserId,
				})
			},
		},

		{
			name:     "サムネイル画像を指定した場合、講習情報の更新と画像アップロードができる",
			courseId: 1,
			requestBody: fmt.Sprintf(`
{
	"title": "ロジカルシンキング研修",
	"thumbnail": {"name": "thumbnail_200_200.png", "content": "%s"},
	"isRemoveThumbnailImage": true,
	"isRequired": true
}`, helper.GetBase64Image(t, "thumbnail_200_200.png")),
			statusCode: http.StatusNoContent,
			after: func(t *testing.T) {
				// DBに保存されたデータを検証
				db := helper.OpenDb(t)
				defer helper.CloseDb(t, db)

				// 講習の値が等しいこと
				assert.EqualFirstRecord(t, db.Where("id = 1"), entity.Course{
					Id:                        1,
					CourseType:                1,
					Title:                     "ロジカルシンキング研修",
					Description:               nil,
					ThumbnailDeliveryFileName: helper.P("thumbnail_200_200-png-hashed.png"),
					ThumbnailOriginalFileName: helper.P("thumbnail_200_200.png"),
					IsRequired:                true,
					CategoryId:                nil,
					Status:                    1,
					CreatedAt:                 helper.FixedTime,
					CreatedBy:                 1,
					UpdatedAt:                 helper.FixedMockTime,
					UpdatedBy:                 helper.TestRequestDefaultUserId,
				})

				// ファイルが存在することを検証
				path := helper.GetTestStorageFilePath("1/thumbnail_200_200-png-hashed.png")
				assert.FileExist(t, path)
			},
		},
		{
			name:     "サムネイル画像を指定せずに、画像削除フラグをONにした場合、講習情報のサムネイル情報がNULLに更新される",
			courseId: 1,
			requestBody: `
{
	"title": "ロジカルシンキング研修",
	"thumbnail": null,
	"isRemoveThumbnailImage": true,
	"isRequired": true
}`,
			statusCode: http.StatusNoContent,
			after: func(t *testing.T) {
				// DBに保存されたデータを検証
				db := helper.OpenDb(t)
				defer helper.CloseDb(t, db)

				// 講習の値が等しいこと
				assert.EqualFirstRecord(t, db.Where("id = 1"), entity.Course{
					Id:                        1,
					CourseType:                1,
					Title:                     "ロジカルシンキング研修",
					Description:               nil,
					ThumbnailDeliveryFileName: nil,
					ThumbnailOriginalFileName: nil,
					IsRequired:                true,
					CategoryId:                nil,
					Status:                    1,
					CreatedAt:                 helper.FixedTime,
					CreatedBy:                 1,
					UpdatedAt:                 helper.FixedMockTime,
					UpdatedBy:                 helper.TestRequestDefaultUserId,
				})
			},
		},

		{
			name:        "ファイルの拡張子がpngの場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.png", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_200_200.png")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "ファイルの拡張子がPNGの場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.PNG", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail.PNG")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "ファイルの拡張子がjpgの場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.jpg", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_200_200.jpg")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "ファイルの拡張子がJPGの場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.JPG", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail.JPG")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "ファイルの拡張子がjpegの場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.jpeg", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_200_200.jpeg")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "ファイルの拡張子がJPEGの場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.JPEG", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail.JPEG")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "ファイルの拡張子がgifの場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.gif", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_200_200.gif")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:        "ファイルの拡張子がGIFの場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.GIF", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail.GIF")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:         "ファイルの拡張子がsvgの場合、422エラーが返却される",
			courseId:     1,
			requestBody:  fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.svg", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail.svg")),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["許可されていない画像の拡張子が指定されています"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "ファイルの拡張子がpngで、実態がsvgの場合、422エラーが返却される",
			courseId:     1,
			requestBody:  fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.png", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail.svg")),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["不正な画像形式のファイルです"]}`,
			after:        func(t *testing.T) {},
		},

		{
			name:         "横幅が200px未満の画像を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.png", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_199_200.jpeg")),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["縦横が200px ~ 2000pxの画像を指定してください"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "縦幅が200px未満の画像を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.png", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_200_199.jpeg")),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["縦横が200px ~ 2000pxの画像を指定してください"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:        "縦横幅が2000pxの画像を指定した場合、講習情報の更新ができる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.png", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_2000_2000.jpeg")),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:         "横幅が2000pxを超える画像を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.png", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_2001_2000.jpeg")),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["縦横が200px ~ 2000pxの画像を指定してください"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "縦幅が2000pxを超える画像を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.png", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_2000_2001.jpeg")),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["縦横が200px ~ 2000pxの画像を指定してください"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "上限値を超える画像を指定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  fmt.Sprintf(`{"title": "ロジカルシンキング研修", "thumbnail": {"name": "thumbnail.png", "content": "%s"}}`, helper.GetBase64Image(t, "thumbnail_large.png")),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["1MB以内の画像を指定してください"]}`,
			after:        func(t *testing.T) {},
		},

		{
			name:         "タイトルを送信しない場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["講習タイトルは必須です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "タイトルに空文字列を設定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": ""}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["講習タイトルは必須です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:        "タイトルに50文字の文字列を設定した場合、講習情報が更新できる",
			courseId:    1,
			requestBody: `{"title": "aAあ🫠漢aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}`,
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:         "タイトルに51文字の文字列を設定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "aAあ🫠漢aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["講習タイトルは50文字以内で入力してください"]}`,
			after:        func(t *testing.T) {},
		},

		{
			name:        "説明文に1000文字の文字列を設定した場合、講習情報が更新できる",
			courseId:    1,
			requestBody: fmt.Sprintf(`{"title": "test", "description": "aAあ🫠漢%s"}`, strings.Repeat("a", 995)),
			statusCode:  http.StatusNoContent,
			after:       func(t *testing.T) {},
		},
		{
			name:         "説明文に1001文字の文字列を設定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  fmt.Sprintf(`{"title": "test", "description": "aAあ🫠漢%s"}`, strings.Repeat("a", 996)),
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["講習の説明は1000文字以内で入力してください"]}`,
			after:        func(t *testing.T) {},
		},

		{
			name:         "存在しないカテゴリIDを設定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "test", "categoryId": 999}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["カテゴリが存在しません"]}`,
			after:        func(t *testing.T) {},
		},

		{
			name:         "期間（開始/終了）に不正な日付を設定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "test", "from": "curry", "to": "pasta"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["期間（開始）の日付形式が不正です", "期間（終了）の日付形式が不正です"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "期間（開始）に期間（終了）と同じ値を設定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "test", "from": "2023/02/20 10:30", "to": "2023/02/20 10:30"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["期間（開始）は期間（終了）の過去日時を指定してください"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "期間（開始）に期間（終了）より未来日を設定した場合、422エラーが返却される",
			courseId:     1,
			requestBody:  `{"title": "test", "from": "2023/02/20 10:31", "to": "2023/02/20 10:30"}`,
			statusCode:   http.StatusUnprocessableEntity,
			expectedBody: `{"errors": ["期間（開始）は期間（終了）の過去日時を指定してください"]}`,
			after:        func(t *testing.T) {},
		},

		{
			name:         "存在しない講習IDを指定した場合、404エラーが返却される",
			courseId:     999,
			requestBody:  `{"title": "e-Learning 1(updated)"}`,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["講習が存在しません"]}`,
			after:        func(t *testing.T) {},
		},
		{
			name:         "講習区分がe-Learning以外の講習IDを指定した場合、404エラーが返却される",
			courseId:     3,
			requestBody:  `{"title": "e-Learning 1(updated)"}`,
			statusCode:   http.StatusNotFound,
			expectedBody: `{"errors": ["講習が存在しません"]}`,
			after:        func(t *testing.T) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			defer helper.CleanTestStorage(t)

			res := helper.DoRequest(t, helper.ApiRequest{
				Method: http.MethodPatch,
				Path:   fmt.Sprintf("/settings/%d/e_learning", tc.courseId),
				Body:   tc.requestBody,
			})
			assert.Equal(tt, res.StatusCode, tc.statusCode)
			assert.EqualJson(tt, string(res.Body), tc.expectedBody)
			tc.after(tt)
		})
	}
}
