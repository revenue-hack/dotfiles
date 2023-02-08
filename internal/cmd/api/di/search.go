//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/search/usecase"
	h "gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
)

func InitializeSearchRequiredHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.New,
		usecase.NewSearchRequired,
		database.NewSearchRequiredQuery,
	)
	return nil
}

func InitializeSearchOptionalHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.New,
		usecase.NewSearchOptional,
		database.NewSearchOptionalQuery,
	)
	return nil
}

func InitializeSearchExpiredHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.New,
		usecase.NewSearchExpired,
		database.NewSearchExpiredQuery,
	)
	return nil
}

func InitializeSearchCompletedHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.New,
		usecase.NewSearchCompleted,
		database.NewSearchCompletedQuery,
	)
	return nil
}
