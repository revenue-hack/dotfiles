//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/handler"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/infrastructure/database"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/service"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/course/usecase"
	h "gitlab.kaonavi.jp/ae/sardine/internal/core/handler"
)

func InitializeSettingGetELearningHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.NewGetELearning,
		usecase.NewGetELearning,
		database.NewGetQuery,
	)
	return nil
}

func InitializeSettingUpdateELearningHandler() h.API {
	wire.Build(
		ProviderSet,
		handler.NewUpdateELearning,
		usecase.NewUpdateELearning,
		database.NewUpdateELearningQuery,
		database.NewUpdateELearningRepository,
		service.NewUpdateELearning,
	)
	return nil
}
