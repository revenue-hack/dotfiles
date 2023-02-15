package content

import (
	"net/http"
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
	"gitlab.kaonavi.jp/ae/sardine/test/helper/assert"
)

func TestSetting_ListContent(t *testing.T) {
	helper.InitDb(t)
	helper.ExecSeeder(t, "setting/content/list")

	expected := `
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
}`

	res := helper.DoRequest(t, helper.ApiRequest{
		Method: http.MethodGet,
		Path:   "/settings/1/contents",
	})
	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.EqualJson(t, string(res.Body), expected)
}
