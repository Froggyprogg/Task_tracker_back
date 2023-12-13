package models

type User struct {
	ID        uint32 `gorm:"primary key;autoIncrement"`
	Login     string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"not null"`
	IsManager bool   `gorm:"not null"`
}

func NewUser(login, password, email string, IsManager bool) *User {
	return &User{Login: login, Password: password, Email: email, IsManager: IsManager}
}
