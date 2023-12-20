package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Report struct {
	gorm.Model
	ReportText string `gorm:"type:varchar(300);not null"`
}

func NewReport(ReportText string) *Report {
	return &Report{ReportText: ReportText}
}

func (r *Report) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("report_id = ?", r.ID).Delete(&Report{})
	return
}
