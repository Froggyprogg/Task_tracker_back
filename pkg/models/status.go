package models

type Status struct {
	IdStatus uint32 `gorm:"primary key;autoIncrement"`
	Name     string `gorm:"type:varchar(100);not null"`
}

func NewStatus(Name string) *Status {
	return &Status{Name: Name}
}
