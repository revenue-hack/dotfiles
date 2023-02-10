//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/course/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/course/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/course/usecase"
	h "gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
)

func InitializeCreateELearningHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.NewCreateHandler,
		usecase.NewCreateELearning,
		database.NewCreateRepository,
	)
	return nil
}
