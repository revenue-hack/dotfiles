package helper

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SetUp() {
	// ginのデバッグログが出て結果が見にくいので、非表示にするためにリリースモードで動かす
	os.Setenv("APP_MODE", gin.ReleaseMode)

	// DBの構築
	if err := createTestDB(dbName); err != nil {
		panic(fmt.Sprintf("SetUp failed to createTestDB %v", err))
	}
}

func TearDown() {
	panicErr := recover()
	err := cleanTestDB()
	if panicErr != nil || err != nil {
		panic(fmt.Sprintf("TearDown revocer: %v, clean: %v", panicErr, err))
	}
}
