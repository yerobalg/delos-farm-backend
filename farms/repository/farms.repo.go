package repository

import (
	"delos-farm-backend/domains"
	"gorm.io/gorm"
)

type FarmsRepository struct {
	conn *gorm.DB
}

//Constructor for farms repository
func NewFarmsRepository(db *gorm.DB) domains.FarmsRepository {
	return &FarmsRepository{db}
}

//Ceate new farm repository
func (r *FarmsRepository) Create(farm *domains.Farms) error{
	return r.conn.Create(farm).Error
}

//delete farm repository
func (r *FarmsRepository) Delete(farm *domains.Farms) error{
	return r.conn.Delete(farm).Error
}
