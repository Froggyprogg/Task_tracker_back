package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Task struct {
	gorm.Model
	Name        string              `gorm:"type:varchar(50);not null"`
	StartDate   timestamp.Timestamp `gorm:"type:date; autoCreateTime;"`
	EndDate     timestamp.Timestamp `gorm:"type:date; not null"`
	Priority    uint                `gorm:"not null"`
	Description string              `gorm:"type:varchar(300); not null"`
	StatusID    uint32              `gorm:"not null"`
	ColumnID    uint32              `gorm:"not null"`
	Status      Status
	Column      Column
}

func NewTask(Name, Description string, Priority uint, ColumnID uint32) *Task {
	return &Task{Name: Name, Description: Description, Priority: Priority, ColumnID: ColumnID}
}

func (t *Task) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("task_id = ?", t.ID).Delete(&Task{})
	return
}
