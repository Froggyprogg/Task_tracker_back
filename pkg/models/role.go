package models

import "gorm.io/gorm"

type Role struct {
	//Id   uint32 `gorm:"primaryKey"`
	gorm.Model
	Name string `gorm:"type:varchar(50);not null"`
}

func NewRole(Name string) *Role {
	return &Role{Name: Name}
}
