package repository

import (
	"delos-farm-backend/domains"
	"gorm.io/gorm"
)

type PondsRepository struct {
	conn *gorm.DB
}

//Constructor for ponds repository
func NewPondsRepository(conn *gorm.DB) domains.PondsRepository {
	return &PondsRepository{conn: conn}
}

//Ceate new pond repository
func (r *PondsRepository) Create(pond *domains.Ponds) error{
	return r.conn.Create(pond).Error
}



