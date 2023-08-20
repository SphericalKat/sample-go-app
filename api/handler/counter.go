package handler

import (
	"context"
	"sample/api/gen"
	"sample/pkg/domain/service"
)

type CounterHandler interface {
	GetCurrentCount(ctx context.Context, request gen.GetCurrentCountRequestObject) (gen.GetCurrentCountResponseObject, error)
}

type counterHandlerImpl struct {
	counterService service.CounterService
}

func NewCounterHandlerImpl(counterService service.CounterService) CounterHandler {
	return &counterHandlerImpl{counterService: counterService}
}

func (ch *counterHandlerImpl) GetCurrentCount(ctx context.Context, request gen.GetCurrentCountRequestObject) (gen.GetCurrentCountResponseObject, error) {
	resp := gen.GetCurrentCount200JSONResponse(gen.CounterResponse{
		Count: 2,
	})
	return resp, nil
}
