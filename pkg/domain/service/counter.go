package service

import "sample/pkg/domain/repository"

type CounterService interface {
	GetCount() (int, error)
	Increment(amount *int) error
	Decrement(amount *int) error
}

type counterServiceImpl struct {
	repo repository.CounterRepository
}

func NewCounterService(repo repository.CounterRepository) CounterService {
	return &counterServiceImpl{repo: repo}
}

func (cs *counterServiceImpl) GetCount() (int, error) {
	panic("todo")
}

func (cs *counterServiceImpl) Increment(amount *int) error {
	panic("todo")
}

func (cs *counterServiceImpl) Decrement(amount *int) error {
	panic("todo")
}
