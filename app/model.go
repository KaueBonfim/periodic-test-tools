package app

import (
	"github.com/jinzhu/gorm"
)
type Tool struct{
		gorm.Model
		Name string `gorm:"type:varchar(100);unique_index" json:"name"`
		Language string `"gorm:"not_null" json:"language"`
		Documentation string `gorm:"type:varchar(100)" json:"documentation"`
		Description string `gorm:"type:varchar(300)" json:"description"`
}



func (Tool) TableName() string {
	return "Tools"
  }