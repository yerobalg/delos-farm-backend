package mocks

import (
	"delos-farm-backend/domains"
	"github.com/stretchr/testify/mock"
)

type PondsServiceMock struct {
	Mock mock.Mock
}

func (s *PondsServiceMock) Create(
	name string, 
	slug string,
	farmId uint,
) (*domains.Ponds, error) {
	args := s.Mock.Called(name, slug, farmId)
	if args.Get(0) == nil && args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).(*domains.Ponds), nil
}

func (s *PondsServiceMock) Delete(pond *domains.Ponds) error {
	args := s.Mock.Called(pond)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (s *PondsServiceMock) Update(pond *domains.Ponds) error {
	args := s.Mock.Called(pond)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (s *PondsServiceMock) Get(id uint) (domains.Ponds, error) {
	args := s.Mock.Called(id)
	if args.Get(0) == nil && args.Get(1) != nil {
		return domains.Ponds{}, args.Get(1).(error)
	}

	ponds := args.Get(0).(domains.Ponds)
	return ponds, nil
}

func (s *PondsServiceMock) GetAll(
	limitInput string,
	offsetInput string,
) ([]domains.Ponds, error) {
	args := s.Mock.Called(limitInput, offsetInput)
	ponds := args.Get(0).([]domains.Ponds)

	if len(ponds) == 0 {
		return []domains.Ponds{}, nil
	}

	return ponds, nil
}
