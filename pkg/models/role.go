package models

type Role struct {
	IdRole uint32 `gorm:"primaryKey;autoIncrement:true"`
	Name   string `gorm:"type:varchar(50);not null"`
}

func NewRole(Name string) *Role {
	return &Role{Name: Name}
}
