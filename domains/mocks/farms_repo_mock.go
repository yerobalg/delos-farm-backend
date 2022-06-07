package mocks

import (
	"delos-farm-backend/domains"
	"github.com/stretchr/testify/mock"
)

type FarmsRepositoryMock struct {
	Mock mock.Mock
}

func (r *FarmsRepositoryMock) Get(id uint) (domains.Farms, error) {
	args := r.Mock.Called(id)
	if args.Get(0) == nil && args.Get(1) != nil {
		return domains.Farms{}, args.Get(1).(error)
	}

	farms := args.Get(0).(domains.Farms)
	return farms, nil
}

func (r *FarmsRepositoryMock) Create(farm *domains.Farms) error {
	args := r.Mock.Called(farm)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *FarmsRepositoryMock) Delete(farm *domains.Farms) error {
	args := r.Mock.Called(farm)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *FarmsRepositoryMock) Update(farm *domains.Farms) error {
	args := r.Mock.Called(farm)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *FarmsRepositoryMock) GetAll(
	limit int,
	offset int,
) ([]domains.Farms, error) {
	args := r.Mock.Called(limit, offset)
	farms := args.Get(0).([]domains.Farms)

	if len(farms) == 0 {
		return []domains.Farms{}, nil
	}

	return farms, nil
}
