package service

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/domains/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var statsRepository = new(mocks.StatsRepositoryMock)
var statsService = NewStatsService(statsRepository)

func TestStatsService_CreateSuccess(t *testing.T) {
	stats := domains.Stats{
		Path: "test",
		IP:   "10.0.0.0",
	}
	statsRepository.Mock.On("CreateStats", stats).Return(nil)

	err := statsService.CreateStats(stats.Path, stats.IP)

	assert.Nil(t, err, "should not return error")
}

func TestStatsService_GetAllSuccess(t *testing.T) {

	statsRepository.Mock.On("GetAllStats", 10, 0).Return(
		[]domains.StatsResults{
			{Path: "test 1"},
			{Path: "test 2"},
		},
		nil,
	)

	statsRes, err := statsService.GetAllStats("10", "0")

	assert.Nil(t, err, "should not return error")
	assert.NotNil(t, statsRes, "stats should exist")
	assert.Equal(t, 2, len(statsRes), "should return 2 stats")
}
