package models

type Board struct {
	IdBoard  uint32 `gorm:"primaryKey;autoIncrement:true"`
	Name     string `gorm:"type:varchar(100);not null"`
	Private  bool   `gorm:"not null; default: false"`
	IdColumn uint32 `gorm:"not null"`
}

func NewBoard(Name string) *Board {
	return &Board{Name: Name}
}
