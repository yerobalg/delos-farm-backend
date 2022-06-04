package domains

import (
	"gorm.io/gorm"
)

type Farms struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name" gorm:"type:varchar(255);column:name;not null"`
	Slug      string         `json:"slug" gorm:"type:varchar(255);column:slug;unique;not null"`
	Ponds     []Ponds        `json:"ponds" gorm:"foreignkey:FarmID;constraint:OnDelete:CASCADE"`
}

type FarmsInput struct {
	Name string `json:"name" binding:"required"`
}

type FarmsService interface {
	Create(farm *Farms) error
	Delete(farm *Farms) error
	Update(farm *Farms) error
	Get(id uint) (Farms, error)
	GetAll(limitInput string, offsetInput string) ([]Farms, error)
}

type FarmsRepository interface {
	Create(farm *Farms) error
	Delete(farm *Farms) error
	Update(farm *Farms) error
	Get(id uint) (Farms, error)
	GetAll(limit int, offset int) ([]Farms, error)
}
