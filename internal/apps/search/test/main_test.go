package search_test

import (
	"testing"

	"gitlab.kaonavi.jp/ae/sardine/test/helper"
)

func TestMain(m *testing.M) {
	defer helper.TearDown()
	helper.SetUp()
	m.Run()
}
