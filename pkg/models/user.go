package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	//Id        uint32 `gorm:"primaryKey"`
	gorm.Model
	Login     string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100); unique; not null"`
	IsManager bool   `gorm:"not null"`
}

func NewUser(login, password, email string, IsManager bool) *User {
	return &User{Login: login, Password: password, Email: email, IsManager: IsManager}
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("user_id = ?", u.ID).Delete(&User{})
	return
}
