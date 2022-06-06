package repository

import (
	"delos-farm-backend/domains"
	"github.com/stretchr/testify/mock"
	"errors"
)

type FarmsRepositoryMock struct {
	Mock mock.Mock
}

func (r *FarmsRepositoryMock) Get(id uint) (domains.Farms, error) {
	args := r.Mock.Called(id)
	if(args.Get(0) == nil && args.Get(1) != nil) {
		return domains.Farms{}, errors.New("Farm not found")
	}

	farms := args.Get(0).(domains.Farms)
	return farms, nil
}

func (r *FarmsRepositoryMock) Create(farms *domains.Farms)  error {
	args := r.Mock.Called(farms)
	if(args.Get(0) == nil) {
		return errors.New("Parameter farms is required")
	}

	return nil
}

func (r *FarmsRepositoryMock) Delete(farms *domains.Farms)  error {
	panic("implement me")
}

func (r *FarmsRepositoryMock) Update(farms *domains.Farms)  error {
	panic("implement me")
}

func (r *FarmsRepositoryMock) GetAll(limit int, offset int) ([]domains.Farms, error) {
	panic("implement me")
}
