//go:generate wire
//go:build wireinject
// +build wireinject

package di

import (
	"sample/api"
	"sample/api/handler"
	"sample/internal/config"
	"sample/pkg/domain/repository"
	"sample/pkg/domain/service"

	"github.com/google/wire"
)

func InjectServer() (*api.ServerImpl, error) {
	wire.Build(
		config.ProvideConfig,
		config.ProvideSqliteDB,
		repository.NewSQLCounterRepository,
		service.NewCounterService,
		handler.NewCounterHandlerImpl,
		api.NewServerImpl,
	)

	return &api.ServerImpl{}, nil
}
