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

//get farm by id repository
func (r *FarmsRepository) Get(id uint) (*domains.Farms, error) {
	var farm domains.Farms
	err := r.conn.First(&farm, id).Error
	return &farm, err
}

 //get all farms repository
func (r *FarmsRepository) GetAll(
	limit int, 
	offset int,
) ([]domains.Farms, int64, error) {
	var farms []domains.Farms
	var totalData int64
	
	if err := r.conn.Limit(limit).Offset(offset).Find(&farms).Error; 
	err != nil {
		return nil, -1, err
	}

	r.conn.Model(&domains.Farms{}).Count(&totalData)
	return farms, totalData, nil
}
