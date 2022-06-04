package domains

import (
	"gorm.io/gorm"
)

type Farms struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);column:name;not null"`
	Slug string `json:"slug" gorm:"type:varchar(255);column:slug;unique;not null"`
	Ponds []Ponds `json:"ponds" gorm:"foreignkey:FarmID;constraint:OnDelete:CASCADE"`
}

type FarmsService interface {
	Create(farm *Farms) error
	Delete(farm *Farms) error
	Get(id uint) (Farms, error)
	GetAll(limitInput string, offsetInput string) ([]Farms, error)
}

type FarmsRepository interface {
	Create(farm *Farms) error
	Delete(farm *Farms) error
	Get(id uint) (Farms, error)
	GetAll(limit int, offset int) ([]Farms, error)
}

