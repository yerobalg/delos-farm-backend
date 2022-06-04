package service

import (
	"delos-farm-backend/domains"
	"strings"
	"errors"
	"strconv"
)

type PondsService struct {
	repo domains.PondsRepository
}

//Constructor for ponds service
func NewPondsService(repo domains.PondsRepository) domains.PondsService {
	return &PondsService{repo: repo}
}

//Create new pond service
func (s *PondsService) Create(pond *domains.Ponds) error {
	err := s.repo.Create(pond)
	if err == nil {
		return nil
	}
	if (strings.Contains(err.Error(), "duplicate key value")) {
		return errors.New("Pond already exists")
	}
	return err	
}

//Delete pond service
func (s* PondsService) Delete(pond *domains.Ponds) error {
	return s.repo.Delete(pond)
}

//Update pond service
func (s *PondsService) Update(pond *domains.Ponds) error {
	err := s.repo.Update(pond)
	if err == nil {
		return nil
	}
	if (strings.Contains(err.Error(), "duplicate key value")) {
		return errors.New("Pond already exists")
	}
	return err	
}

//Get pond by id service
func (s *PondsService) Get(id uint) (domains.Ponds, error) {
	pond, err := s.repo.Get(id)
	if err == nil {
		return pond, nil
	}

	if (strings.Contains(err.Error(), "record not found")) {
		return pond, errors.New("Pond not found")
	}

	return pond, err
}

//Get all ponds service
func (s *PondsService) GetAll(
	limitInput string,
	offsetInput string,
) ([]domains.Ponds, error) {
	limit, _ := strconv.Atoi(limitInput)
	offset, _ := strconv.Atoi(offsetInput)
	
	ponds, err := s.repo.GetAll(limit, offset)

	if err == nil && len(ponds) == 0 { 
		return ponds, errors.New("No ponds found")
	}
	return ponds, err
}



