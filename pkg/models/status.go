package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`
}

func NewStatus(Name string) *Status {
	return &Status{Name: Name}
}
