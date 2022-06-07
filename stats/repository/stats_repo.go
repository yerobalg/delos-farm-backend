package repository

import (
	"delos-farm-backend/domains"
	"gorm.io/gorm"
)

type StatsRepository struct {
	conn *gorm.DB
}

//Constructor for stats repository
func NewStatsRepository(conn *gorm.DB) domains.StatsRepository {
	return &StatsRepository{conn: conn}
}

//Create stats
func (r *StatsRepository) CreateStats(stats domains.Stats) error {
	return r.conn.Create(&stats).Error
}

//Get all stats
func (r *StatsRepository) GetAllStats(
	limit int,
	offset int,
) ([]domains.StatsResults, error) {
	var results []domains.StatsResults
	err := r.conn.Model(&domains.Stats{}).Select(
		"path",
		"count(ip) as api_call_count",
		"count(distinct ip) as unique_call_count",
	).Group("path").Limit(limit).Offset(offset).Find(&results).Error

	return results, err
}
