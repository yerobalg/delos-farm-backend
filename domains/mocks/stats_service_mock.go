package mocks

import (
	"delos-farm-backend/domains"
	"github.com/stretchr/testify/mock"
)

type StatsServiceMock struct {
	Mock mock.Mock
}

// Mock for CreateStats
func (s *StatsServiceMock) CreateStats(path string, ip string) error {
	args := s.Mock.Called(path, ip)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

// Mock for GetAllStats
func (s *StatsServiceMock) GetAllStats(
	limit string,
	offset string,
) ([]domains.StatsResults, error) {
	args := s.Mock.Called(limit, offset)

	if args.Get(1) != nil {
		return []domains.StatsResults{}, args.Get(1).(error)
	}

	return args.Get(0).([]domains.StatsResults), nil
}
