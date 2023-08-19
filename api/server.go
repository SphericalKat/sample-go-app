package api

import "sample/api/handler"

type ServerImpl struct {
	handler.CounterHandler
}

func NewServerImpl(
	counterHandler handler.CounterHandler,
) *ServerImpl {
	return &ServerImpl{
		CounterHandler: counterHandler,
	}
}
