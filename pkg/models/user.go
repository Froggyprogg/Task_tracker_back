package models

type User struct {
	IdUser    uint32 `gorm:"primaryKey;autoIncrement:true"`
	Login     string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(32);not null"`
	Email     string `gorm:"type:varchar(100); not null"`
	IsManager bool   `gorm:"not null"`
}

func NewUser(login, password, email string, IsManager bool) *User {
	return &User{Login: login, Password: password, Email: email, IsManager: IsManager}
}
