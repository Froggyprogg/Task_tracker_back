package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Task struct {
	IdTask      uint32 `gorm:"primaryKey;autoIncrement"`
	IdColumn    uint32
	Name        string `gorm:"type:varchar(50);not null"`
	StartDate   timestamp.Timestamp
	EndDate     timestamp.Timestamp
	Priority    uint32
	Description string
	IdStatus    uint32
}
