package repository

import (
	"context"
	"delos-farm-backend/domains"
	"github.com/go-redis/redis/v8"
)

type StatsRepository struct {
	redisClient *redis.Client
}

//Constructor for ponds repository
func NewStatsRepository(rc *redis.Client) domains.StatsRepository {
	return &StatsRepository{redisClient: rc}
}

//Count api call repository
func (r *StatsRepository) CountAPICall(path string) (int64, error) {
	return r.redisClient.Incr(context.Background(), path).Result()
}

//Count unique call reposit
func (r *StatsRepository) CountUniqueCall(ip string) (int64, error) {
	_, err := r.redisClient.SAdd(context.Background(), "unique_ip", ip).Result()
	if err != nil {
		return -1, err
	}
	return r.redisClient.SCard(context.Background(), "unique_ip").Result()
}
