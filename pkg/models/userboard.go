package models

import "gorm.io/gorm"

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
