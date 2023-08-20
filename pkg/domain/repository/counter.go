package repository

import "sample/ent"

type CounterRepository interface {
	GetCount() (int, error)
	SetCount(count int) error
}

type counterRepositoryImpl struct {
	db *ent.Client
}

func NewSQLCounterRepository(db *ent.Client) CounterRepository {
	return &counterRepositoryImpl{db: db}
}

func (cr *counterRepositoryImpl) GetCount() (int, error) {
	panic("todo")
}

func (cr *counterRepositoryImpl) SetCount(count int) error {
	panic("todo")
}
