//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content/service"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content/usecase"
	h "gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
)

func InitializeSettingListContentHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.NewList,
		usecase.NewList,
		database.NewQuery,
		database.NewListQuery,
	)
	return nil
}

func InitializeSettingUrlCreateHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.NewUrlCreate,
		usecase.NewUrlCreate,
		database.NewQuery,
		database.NewUrlCreateRepository,
		service.NewUrl,
	)
	return nil
}
