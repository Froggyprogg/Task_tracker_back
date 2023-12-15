package models

type Column struct {
	IdColumn int32  `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(100);not null"`
	IdTask   int32  `gorm:"not null"`
}
