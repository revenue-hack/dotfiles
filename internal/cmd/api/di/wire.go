package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
)

var (
	ProviderSet = wire.NewSet(
		database.NewFactory,
	)
)
