package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"gorm.io/gorm"
)

type Subtask struct {
	gorm.Model
	Name        string              `gorm:"type:varchar(50);not null"`
	StartDate   timestamp.Timestamp `gorm:"type:date; autoCreateTime;"`
	EndDate     timestamp.Timestamp `gorm:"type:date; not null"`
	Priority    uint                `gorm:"not null"`
	Description string              `gorm:"type:varchar(300); not null"`
	StatusID    uint32              `gorm:"not null"`
	TaskID      uint32              `gorm:"not null"`
	Status      Status
	Task        Task
}

/*func NewSubtask(Name, Description string, IdStatus, IdTask uint32, Priority bool, EndDate timestamp.Timestamp) *Subtask {
	return &Subtask{Name: Name, Description: Description, IdStatus: IdStatus, IdTask: IdTask, Priority: Priority, EndDate: EndDate}
}*/
