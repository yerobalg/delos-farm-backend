package mocks

import (
	"delos-farm-backend/domains"
	"github.com/stretchr/testify/mock"
)

type PondsRepositoryMock struct {
	Mock mock.Mock
}

func (r *PondsRepositoryMock) Get(id uint) (domains.Ponds, error) {
	args := r.Mock.Called(id)
	if args.Get(0) == nil && args.Get(1) != nil {
		return domains.Ponds{}, args.Get(1).(error)
	}

	pond := args.Get(0).(domains.Ponds)
	return pond, nil
}

func (r *PondsRepositoryMock) Create(pond *domains.Ponds) error {
	args := r.Mock.Called(pond)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *PondsRepositoryMock) Delete(pond *domains.Ponds) error {
	args := r.Mock.Called(pond)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *PondsRepositoryMock) Update(pond *domains.Ponds) error {
	args := r.Mock.Called(pond)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *PondsRepositoryMock) GetAll(
	limit int,
	offset int,
) ([]domains.Ponds, error) {
	args := r.Mock.Called(limit, offset)
	ponds := args.Get(0).([]domains.Ponds)

	if len(ponds) == 0 {
		return []domains.Ponds{}, nil
	}

	return ponds, nil
}
