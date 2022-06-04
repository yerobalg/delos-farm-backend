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
func NewpondsService(repo domains.PondsRepository) domains.PondsService {
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
