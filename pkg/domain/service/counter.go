package service

import (
	"context"
	"errors"
	"fmt"
	"sample/pkg/domain/repository"
)

type CounterService interface {
	GetCount(ctx context.Context) (int, error)
	Increment(ctx context.Context, amount *int) error
	Decrement(ctx context.Context, amount *int) error
}

type counterServiceImpl struct {
	repo repository.CounterRepository
}

func NewCounterService(repo repository.CounterRepository) CounterService {
	return &counterServiceImpl{repo: repo}
}

func (cs *counterServiceImpl) GetCount(ctx context.Context) (int, error) {
	return cs.repo.GetCount(ctx)
}

func (cs *counterServiceImpl) Increment(ctx context.Context, amount *int) error {
	currentCount, err := cs.repo.GetCount(ctx)
	if err != nil {
		return err
	}

	incrementAmount := 1
	if amount != nil {
		incrementAmount = *amount
	}

	return cs.repo.SetCount(ctx, currentCount + incrementAmount)
}

func (cs *counterServiceImpl) Decrement(ctx context.Context, amount *int) error {
	currentCount, err := cs.repo.GetCount(ctx)
	if err != nil {
		return err
	}

	incrementAmount := 1
	if amount != nil {
		incrementAmount = *amount
	}

	if currentCount - incrementAmount < 0 {
		return fmt.Errorf("error decrementing count: %w", errors.New("negative counts are not allowed"))
	}

	return cs.repo.SetCount(ctx, currentCount - incrementAmount)
}
