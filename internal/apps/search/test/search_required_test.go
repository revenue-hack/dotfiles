package search_test

import (
	"fmt"
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSearchRequired(t *testing.T) {
	helper.InitDb(t, "testdata/testdata.sql")

	res := helper.DoRequest(t, helper.ApiRequest{
		Method: http.MethodPost,
		Path:   "/search/required",
		Body:   `{}`,
	})

	expected := fmt.Sprintf(`
{
	"maxPageSize": 0,
	"nextPageToken": null,
	"courses": [
		{
			"id": 1,
			"title": "e-Learning 1",
			"thumbnail": {"name": "original1.png", "url": "%s"},
			"categoryName": null,
			"expireAt": "2023/03/01 18:30:00",
			"recommend": null,
			"isRequired": false,
			"isFixed": false
		},
		{
			"id": 2,
			"title": "e-Learning 2",
			"thumbnail": null,
			"categoryName": null,
			"expireAt": "2023/03/02 18:30:00",
			"recommend": null,
			"isRequired": false,
			"isFixed": false
		},
		{
			"id": 3,
			"title": "e-Learning 3",
			"thumbnail": null,
			"categoryName": null,
			"expireAt": null,
			"recommend": null,
			"isRequired": false,
			"isFixed": false
		},
		{
			"id": 4,
			"title": "e-Learning 4",
			"thumbnail": null,
			"categoryName": null,
			"expireAt": null,
			"recommend": null,
			"isRequired": false,
			"isFixed": false
		},
		{
			"id": 5,
			"title": "e-Learning 5",
			"thumbnail": {"name": "original5.png", "url": "%s"},
			"categoryName": null,
			"expireAt": "2023/03/05 18:30:00",
			"recommend": null,
			"isRequired": true,
			"isFixed": false
		},
		{
			"id": 6,
			"title": "e-Learning 6",
			"thumbnail": null,
			"categoryName": null,
			"expireAt": "2023/03/06 18:30:00",
			"recommend": null,
			"isRequired": true,
			"isFixed": false
		},
		{
			"id": 7,
			"title": "e-Learning 7",
			"thumbnail": null,
			"categoryName": null,
			"expireAt": null,
			"recommend": null,
			"isRequired": true,
			"isFixed": false
		},
		{
			"id": 8,
			"title": "e-Learning 8",
			"thumbnail": null,
			"categoryName": null,
			"expireAt": null,
			"recommend": null,
			"isRequired": true,
			"isFixed": false
		}
	]
}`,
		helper.GetThumbnailDeliveryUrl(1, "delivery1.png"),
		helper.GetThumbnailDeliveryUrl(5, "delivery5.png"),
	)

	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.EqualJson(t, string(res.Body), expected)
}
