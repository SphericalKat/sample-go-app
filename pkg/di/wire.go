//go:generate wire
//go:build wireinject
// +build wireinject

package di

import (
	"sample/api"
	"sample/api/handler"

	"github.com/google/wire"
)

func InjectServer() (*api.ServerImpl, error) {
	wire.Build(
		handler.NewCounterHandlerImpl,
		api.NewServerImpl,
	)

	return &api.ServerImpl{}, nil
}
