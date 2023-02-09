package search

import (
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSearchRequired(t *testing.T) {
	helper.InitDb(t)
	helper.ExecSeeder(t, "search/search")

	res := helper.DoRequest(t, helper.ApiRequest{
		Method: http.MethodPost,
		Path:   "/search/required",
		Body:   `{}`,
	})

	expected := `
{
	"maxPageSize": 0,
	"nextPageToken": null,
	"courses": [
		{
			"id": 1,
			"title": "e-Learning 1",
			"thumbnailUrl": "",
			"categoryName": null,
			"expireAt": "2023/03/01 18:30:00",
			"recommend": null,
			"isRequired": false,
			"isFixed": false
		},
		{
			"id": 2,
			"title": "e-Learning 2",
			"thumbnailUrl": "",
			"categoryName": null,
			"expireAt": "2023/03/02 18:30:00",
			"recommend": null,
			"isRequired": false,
			"isFixed": false
		},
		{
			"id": 3,
			"title": "e-Learning 3",
			"thumbnailUrl": "",
			"categoryName": null,
			"expireAt": null,
			"recommend": null,
			"isRequired": false,
			"isFixed": false
		},
		{
			"id": 4,
			"title": "e-Learning 4",
			"thumbnailUrl": "",
			"categoryName": null,
			"expireAt": null,
			"recommend": null,
			"isRequired": false,
			"isFixed": false
		},
		{
			"id": 5,
			"title": "e-Learning 5",
			"thumbnailUrl": "",
			"categoryName": null,
			"expireAt": "2023/03/05 18:30:00",
			"recommend": null,
			"isRequired": true,
			"isFixed": false
		},
		{
			"id": 6,
			"title": "e-Learning 6",
			"thumbnailUrl": "",
			"categoryName": null,
			"expireAt": "2023/03/06 18:30:00",
			"recommend": null,
			"isRequired": true,
			"isFixed": false
		},
		{
			"id": 7,
			"title": "e-Learning 7",
			"thumbnailUrl": "",
			"categoryName": null,
			"expireAt": null,
			"recommend": null,
			"isRequired": true,
			"isFixed": false
		},
		{
			"id": 8,
			"title": "e-Learning 8",
			"thumbnailUrl": "",
			"categoryName": null,
			"expireAt": null,
			"recommend": null,
			"isRequired": true,
			"isFixed": false
		}
	]
}`

	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.EqualJson(t, string(res.Body), expected)
}
