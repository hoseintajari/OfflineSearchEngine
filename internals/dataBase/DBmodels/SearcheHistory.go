package DBmodels

import "gorm.io/gorm"

type SearchHistory struct {
	gorm.Model
	UserName string `gorm:"ForeignKey:ID" binding:"required"`
	Search   string `gorm:"size:255" json:"password,omitempty" binding:"required"`
}
