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
