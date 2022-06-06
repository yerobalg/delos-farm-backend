package repository

import (
	"github.com/stretchr/testify/mock"
)

type StatsRepositoryMock struct {
	Mock mock.Mock
}

func (r *StatsRepositoryMock) CountAPICall(path string) (int64, error) {
	args := r.Mock.Called(path)
	if args.Get(0) == nil && args.Get(1) != nil {
		return -1, args.Get(1).(error)
	}

	return args.Get(0).(int64), nil
}

func (r *StatsRepositoryMock) CountUniqueCall(ip string) (int64, error) {
	args := r.Mock.Called(ip)
	if args.Get(0) == nil && args.Get(1) != nil {
		return -1, args.Get(1).(error)
	}

	return args.Get(0).(int64), nil
}