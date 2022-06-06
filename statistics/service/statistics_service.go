package service

import (
	"delos-farm-backend/domains"
)

type StatsService struct {
	repo domains.StatsRepository
}

func NewStatsService(repo domains.StatsRepository) domains.StatsService {
	return &StatsService{repo: repo}
}

func (s *StatsService) GetStatistics(path string, ip string) domains.Stats {
	stats := domains.Stats{
		APICount:        0,
		UniqueCallCount: 0,
	}

	apiCount, err := s.repo.CountAPICall(path)
	if err != nil {
		return stats
	}
	stats.APICount = apiCount

	uniqueCallCount, err := s.repo.CountUniqueCall(ip)
	if err != nil {
		return stats
	}
	stats.UniqueCallCount = uniqueCallCount

	return stats
}
