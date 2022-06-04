package service

import (
	"delos-farm-backend/domains"
	"strings"
	"errors"
)

type FarmsService struct {
	repo domains.FarmsRepository
}

func NewFarmsService(repo domains.FarmsRepository) domains.FarmsService {
	return &FarmsService{repo: repo}
}

func (s *FarmsService) Create(farm *domains.Farms) error {
	err := s.repo.Create(farm)
	if err == nil {
		return nil
	}
	if (strings.Contains(err.Error(), "duplicate key value")) {
		return errors.New("Farm already exists")
	}
	return err	
}

func (s* FarmsService) Delete(farm *domains.Farms) error {
	return s.repo.Delete(farm)
}
