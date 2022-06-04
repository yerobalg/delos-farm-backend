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

//update farm repository
func (r* FarmsRepository) Update(farm *domains.Farms) error{
	return r.conn.Save(farm).Error
}

//get farm by id repository
func (r *FarmsRepository) Get(id uint) (domains.Farms, error) {
	var farm domains.Farms
	err := r.conn.First(&farm, id).Error
	return farm, err
}

 //get all farms repository
func (r *FarmsRepository) GetAll() ([]domains.Farms, error) {
	var farms []domains.Farms
	err := r.conn.Find(&farms).Error
	return farms, err
}
