package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Subtask struct {
	IdSubtask   uint32              `gorm:"primary key;autoIncrement"`
	Name        string              `gorm:"type:varchar(50);not null"`
	StartDate   timestamp.Timestamp `gorm:"type:date; autoCreateTime;"`
	EndDate     timestamp.Timestamp `gorm:"type:date; not null"`
	Priority    bool                `gorm:"type:boolean;not null"`
	Description string              `gorm:"type:varchar(300); not null"`
	IdStatus    uint32
	IdTask      uint32
}

func NewSubtask(Name, Description string, IdStatus, IdTask uint32, Priority bool, EndDate timestamp.Timestamp) *Subtask {
	return &Subtask{Name: Name, Description: Description, IdStatus: IdStatus, IdTask: IdTask, Priority: Priority, EndDate: EndDate}
}
