package models

import "gorm.io/gorm"

type Column struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	BoardID uint32 `gorm:"not null"`
	//IdTask int32  `gorm:"not null"`
	//Task   Task   `gorm:"foreignKey:IdTask;references:Id"`
}

func NewColumn(Name string, BoardID uint32) *Column {
	return &Column{Name: Name, BoardID: BoardID}
}
