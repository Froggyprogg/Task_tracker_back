package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	ReportText string `gorm:"type:varchar(300);not null"`
}

func NewReport(ReportText string) *Report {
	return &Report{ReportText: ReportText}
}
