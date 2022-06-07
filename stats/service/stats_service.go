package service

import (
	"delos-farm-backend/domains"
	"strconv"
)

type StatsService struct {
	repo domains.StatsRepository
}

func NewStatsService(repo domains.StatsRepository) domains.StatsService {
	return &StatsService{repo: repo}
}

func (s *StatsService) CreateStats(path string, ip string) error {
	stats := domains.Stats{Path: path, IP: ip}
	return s.repo.CreateStats(stats)
}

func (s *StatsService) GetAllStats(
	limit string, 
	offset string,
) ([]domains.StatsResults, error) {
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	return s.repo.GetAllStats(limitInt, offsetInt)
}