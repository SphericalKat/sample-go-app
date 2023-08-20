package repository

import (
	"context"
	"log/slog"
	"sample/ent"
	"sample/ent/counter"
)

type CounterRepository interface {
	GetCount(ctx context.Context) (int, error)
	SetCount(ctx context.Context, count int) error
}

type counterRepositoryImpl struct {
	db *ent.Client
}

func NewSQLCounterRepository(db *ent.Client) CounterRepository {
	return &counterRepositoryImpl{db: db}
}

func (cr *counterRepositoryImpl) GetCount(ctx context.Context) (int, error) {
	counter, err := cr.db.Counter.Get(ctx, 1)
	if err != nil {
		slog.Error("error fetching count", "err", err)
		return 0, err
	}

	if counter == nil {
		err = cr.db.Counter.Create().SetCount(1).Exec(ctx)
		if err != nil {
			slog.Error("error creating new counter entity", "err", err)
			return 0, err
		}

		return 1, nil
	}

	return counter.Count, nil
}

func (cr *counterRepositoryImpl) SetCount(ctx context.Context, count int) error {
	return cr.db.Counter.Update().Where(counter.IDEQ(1)).SetCount(count).Exec(ctx)
}
