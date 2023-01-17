package main

import (
	"os"

	"gitlab.kaonavi.jp/ae/sardine/internal/cmd/api"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

func main() {
	logger.New(os.Stdout)

	if err := api.Route().Run(); err != nil {
		panic(err)
	}
}
