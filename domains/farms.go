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

type FarmsRepository interface {
	Create(farm *Farms) error
	
}
