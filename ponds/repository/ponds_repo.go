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

//delete pond repository
func (r *PondsRepository) Delete(pond *domains.Ponds) error{
	return r.conn.Delete(pond).Error
}

//update pond repository
func (r* PondsRepository) Update(pond *domains.Ponds) error{
	return r.conn.Save(pond).Error
}

//get pond by id repository
func (r *PondsRepository) Get(id uint) (domains.Ponds, error) {
	var pond domains.Ponds
	err := r.conn.First(&pond, id).Error
	return pond, err
}

