package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserBoard struct {
	gorm.Model
	UserID  uint32 `gorm:"not null"`
	BoardID uint32 `gorm:"not null"`
	RoleID  uint32 `gorm:"not null"`
	User    User
	Board   Board
	Role    Role
}

func NewUserBoard(UserID, BoardID, RoleID uint32) *UserBoard {
	return &UserBoard{UserID: UserID, BoardID: BoardID, RoleID: RoleID}
}

func (u *UserBoard) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("user_board_id = ?", u.ID).Delete(&UserBoard{})
	return
}
