package repository

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

func (r *FarmsRepositoryMock) Create(farms *domains.Farms) error {
	args := r.Mock.Called(farms)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *FarmsRepositoryMock) Delete(farms *domains.Farms) error {
	args := r.Mock.Called(farms)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *FarmsRepositoryMock) Update(farms *domains.Farms) error {
	args := r.Mock.Called(farms)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *FarmsRepositoryMock) GetAll(limit int, offset int) ([]domains.Farms, error) {
	args := r.Mock.Called(limit, offset)
	farms := args.Get(0).([]domains.Farms)

	if args.Get(1) != nil && len(farms) == 0 {
		return nil, args.Get(1).(error)
	}

	return farms, nil
}
