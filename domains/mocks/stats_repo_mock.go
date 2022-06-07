package mocks

import (
	"github.com/stretchr/testify/mock"
	"delos-farm-backend/domains"
)

type StatsRepositoryMock struct {
	Mock mock.Mock
}

// Mock for CreateStats
func (r *StatsRepositoryMock) CreateStats(stats domains.Stats) error {
	args := r.Mock.Called(stats)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}


// Mock for GetAllStats
func (r *StatsRepositoryMock) GetAllStats(
	limit int,
	offset int,
) ([]domains.StatsResults, error) {
	args := r.Mock.Called(limit, offset)
	if args.Get(1) != nil { 
		return []domains.StatsResults{}, args.Get(1).(error)
	}

	return args.Get(0).([]domains.StatsResults), nil
}