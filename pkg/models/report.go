package models

type Report struct {
	IdReport   uint32 `gorm:"primaryKey;autoIncrement"`
	ReportText string `gorm:"type:varchar(300);not null"`
}

func NewReport(ReportText string) *Report {
	return &Report{ReportText: ReportText}
}
