package domains

import (
	"gorm.io/gorm"
)

type Ponds struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	FarmID    uint           `json:"farm_id"`
	Farms     Farms          `json:"farms" gorm:"foreignkey:FarmID"`
}

type PondsInput struct {
	Name string `json:"name" binding:"required"`
}

type PondsService interface {
	Create(pond *Ponds) error
	Delete(pond *Ponds) error
	Update(pond *Ponds) error
	Get(id uint) (Ponds, error)
	GetAll(limitInput string, offsetInput string) ([]Ponds, error)
}

type PondsRepository interface {
	Create(pond *Ponds) error
	Delete(pond *Ponds) error
	Update(pond *Ponds) error
	Get(id uint) (Ponds, error)
	GetAll(limit int, offset int) ([]Ponds, error)
}
