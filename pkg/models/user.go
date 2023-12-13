package models

type User struct {
	ID        uint32 `gorm:"primary key;autoIncrement"`
	Login     string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"not null"`
	IsManager bool   `gorm:"not null"`
}
