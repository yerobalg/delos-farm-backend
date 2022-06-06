package service

import (
	"delos-farm-backend/stats/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var statsRepository = new(repository.StatsRepositoryMock)
var statsService = NewStatsService(statsRepository)

func TestStatsService_GetStatisticsSuccess(t *testing.T) {
	path := "GET_/api/v1/farms"
	ip := "ip_::1"
	statsRepository.Mock.On("CountAPICall", path).Return(int64(1), nil)
	statsRepository.Mock.On("CountUniqueCall", ip).Return(int64(1), nil)

	stats := statsService.GetStatistics(path, ip)

	assert.NotNil(t, stats, "Stats should exist")
	assert.Equal(t, int64(1), stats.APICount, "APICount should be 1")
	assert.Equal(
		t, 
		int64(1), 
		stats.UniqueCallCount, "UniqueCallCount should be 1",
	)
}

