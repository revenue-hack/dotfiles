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
	"thumbnailImage": null,
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
					Id:                 1,
					CourseType:         1,
					Title:              "ロジカルシンキング研修",
					Description:        helper.P("ロジカルな思考をゲット"),
					ThumbnailImageName: nil,
					IsRequired:         true,
					CategoryId:         helper.P[uint32](2),
					Status:             1,
					CreatedAt:          helper.FixedTime,
					CreatedBy:          1,
					UpdatedAt:          helper.FixedMockTime,
					UpdatedBy:          helper.TestRequestDefaultUserId,
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
	"thumbnailImage": null,
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
					Id:                 2,
					CourseType:         1,
					Title:              "ロジカルシンキング研修",
					Description:        helper.P("ロジカルな思考をゲット"),
					ThumbnailImageName: nil,
					IsRequired:         false,
					CategoryId:         helper.P[uint32](1),
					Status:             2,
					CreatedAt:          helper.FixedTime,
					CreatedBy:          1,
					UpdatedAt:          helper.FixedMockTime,
					UpdatedBy:          helper.TestRequestDefaultUserId,
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
