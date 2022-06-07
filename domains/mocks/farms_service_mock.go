package mocks

import (
	"delos-farm-backend/domains"
	"github.com/stretchr/testify/mock"
)

type FarmsServiceMock struct {
	Mock mock.Mock
}

func (s *FarmsServiceMock) Create(
	name string, 
	slug string,
) (*domains.Farms, error) {
	args := s.Mock.Called(name, slug)
	if args.Get(0) == nil && args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).(*domains.Farms), nil
}

func (s *FarmsServiceMock) Delete(farm *domains.Farms) error {
	args := s.Mock.Called(farm)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (s *FarmsServiceMock) Update(farm *domains.Farms) error {
	args := s.Mock.Called(farm)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (s *FarmsServiceMock) Get(id uint) (domains.Farms, error) {
	args := s.Mock.Called(id)
	if args.Get(0) == nil && args.Get(1) != nil {
		return domains.Farms{}, args.Get(1).(error)
	}

	farms := args.Get(0).(domains.Farms)
	return farms, nil
}

func (s *FarmsServiceMock) GetAll(
	limitInput string,
	offsetInput string,
) ([]domains.Farms, error) {
	args := s.Mock.Called(limitInput, offsetInput)
	farms := args.Get(0).([]domains.Farms)

	if len(farms) == 0 {
		return []domains.Farms{}, nil
	}

	return farms, nil
}
