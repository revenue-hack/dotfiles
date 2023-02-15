package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/storage"
)

var (
	ProviderSet = wire.NewSet(
		database.NewFactory,
		storage.NewFactory,
	)
)
