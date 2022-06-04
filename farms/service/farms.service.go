package service

import (
	"delos-farm-backend/domains"
	"strings"
	"errors"
	"strconv"
)

type FarmsService struct {
	repo domains.FarmsRepository
}

//Constructor for farms service
func NewFarmsService(repo domains.FarmsRepository) domains.FarmsService {
	return &FarmsService{repo: repo}
}

//Create new farm service
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

//Delete farm service
func (s* FarmsService) Delete(farm *domains.Farms) error {
	return s.repo.Delete(farm)
}

//Get farm by id service
func (s *FarmsService) Get(id uint) (domains.Farms, error) {
	farm, err := s.repo.Get(id)
	if err == nil {
		return farm, nil
	}

	if (strings.Contains(err.Error(), "record not found")) {
		return farm, errors.New("Farm not found")
	}

	return farm, err
}

 //Get all farms service
 func (s *FarmsService) GetAll(
	 limitInput string,
	 offsetInput string,
	) ([]domains.Farms, error) {
	limit, _ := strconv.Atoi(limitInput)
	offset, _ := strconv.Atoi(offsetInput)

	farms, err := s.repo.GetAll(limit, offset)
	
	if err == nil && len(farms) == 0 { 
		return farms, errors.New("No farms found")
	}
	return farms, err
}
