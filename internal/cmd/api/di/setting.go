//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/infrastructure/storage"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/service"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/usecase"
	h "gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
	coreCdn "gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/cdn"
	coreStorage "gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/storage"
)

func InitializeSettingGetELearningHandler() h.API {
	wire.Build(
		ProviderSet,
		coreCdn.NewFactory,
		handler.NewGetELearning,
		usecase.NewGetELearning,
		database.NewGetQuery,
	)
	return nil
}

func InitializeSettingUpdateELearningHandler() h.API {
	wire.Build(
		ProviderSet,
		coreStorage.NewFactory,
		handler.NewUpdateELearning,
		usecase.NewUpdateELearning,
		database.NewUpdateELearningQuery,
		database.NewUpdateELearningRepository,
		storage.NewThumbnailRepository,
		service.NewUpdateELearning,
	)
	return nil
}
