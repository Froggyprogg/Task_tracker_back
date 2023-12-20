package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserBoardID uint32
	TaskID      uint32
	SubtaskID   uint32
	ReportID    uint32
	UserBoard   UserBoard
	Task        Task
	Subtask     Subtask
	Report      Report
	Comment     string `gorm:"not null"`
}
