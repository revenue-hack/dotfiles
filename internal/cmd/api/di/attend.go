//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/attend/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/attend/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/attend/usecase"
	h "gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	coreCdn "gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/cdn"
)

func InitializeAttendGetELearningHandler() h.API {
	wire.Build(
		ProviderSet,
		coreCdn.NewFactory,
		handler.NewGetELearning,
		usecase.NewGetELearning,
		database.NewGetQuery,
	)
	return nil
}
