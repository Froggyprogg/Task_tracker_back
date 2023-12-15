package models

type UserBoard struct {
	IdUserBoard uint32 `gorm:"primaryKey;autoIncrement:true"`
	IdRole      uint32
	IdBoard     uint32
	IdUser      uint32
	Role        []Role  `gorm:"foreignKey:IdRole;references:IdRole"`
	Board       []Board `gorm:"foreignKey:IdBoard;references:IdBoard"`
	User        []User  `gorm:"foreignKey:IdUser;references:IdUser"`
}

/*type UserBoard struct {
	IdUserBoard uint32 `gorm:"primary key;autoIncrement"`
	IdRole      uint32 `gorm:"not null"`
	Role        Role   `gorm:"foreignKey:IdRole"`
	IdBoard     uint32 `gorm:"not null"`
	Board       Board  `gorm:"foreignKey:IdBoard"`
	IdUser      uint32 `gorm:"not null"`
	User        User   `gorm:"foreignKey:IdUser"`
}*/
