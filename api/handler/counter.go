package handler

import (
	"context"
	"sample/api/gen"
	"sample/pkg/domain/service"
)

type CounterHandler interface {
	GetCurrentCount(ctx context.Context, request gen.GetCurrentCountRequestObject) (gen.GetCurrentCountResponseObject, error)
	IncrementCount(ctx context.Context, request gen.IncrementCountRequestObject) (gen.IncrementCountResponseObject, error)
}

type counterHandlerImpl struct {
	counterService service.CounterService
}

func NewCounterHandlerImpl(counterService service.CounterService) CounterHandler {
	return &counterHandlerImpl{counterService: counterService}
}

func (ch *counterHandlerImpl) GetCurrentCount(ctx context.Context, request gen.GetCurrentCountRequestObject) (gen.GetCurrentCountResponseObject, error) {
	count, err := ch.counterService.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	return gen.GetCurrentCount200JSONResponse(gen.CounterResponse{
		Count: float32(count),
	}), nil
}

func (ch *counterHandlerImpl) IncrementCount(ctx context.Context, request gen.IncrementCountRequestObject) (gen.IncrementCountResponseObject, error) {
	err := ch.counterService.Increment(ctx, request.Body.Amount)
	if err != nil {
		return nil, err
	}

	count, err := ch.counterService.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	return gen.IncrementCount200JSONResponse(gen.CounterResponse{
		Count: float32(count),
	}), nil
}
