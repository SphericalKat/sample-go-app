package handler

import (
	"context"
	"sample/api/gen"
)

type CounterHandler interface {
	GetCurrentCount(ctx context.Context, request gen.GetCurrentCountRequestObject) (gen.GetCurrentCountResponseObject, error)
}

type counterHandlerImpl struct {}

func NewCounterHandlerImpl() CounterHandler {
	return &counterHandlerImpl{}
}

func (ch *counterHandlerImpl) GetCurrentCount(ctx context.Context, request gen.GetCurrentCountRequestObject) (gen.GetCurrentCountResponseObject, error) {
	resp := gen.GetCurrentCount200JSONResponse(gen.CounterResponse{
		Count: 2,
	})
	return resp, nil
}
