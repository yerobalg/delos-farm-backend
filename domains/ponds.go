package domains

import (
	"gorm.io/gorm"
)

type Ponds struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
	FarmID uint `json:"farm_id"`
	Farms Farms `json:"farms" gorm:"foreignkey:FarmID"`
}
