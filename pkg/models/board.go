package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	Private bool   `gorm:"not null; default: false"`
	Column  []Column
}

func NewBoard(Name string, Private bool) *Board {
	return &Board{Name: Name, Private: Private}
}
