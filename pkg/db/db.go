package db

import (
	"fmt"
	"github.com/Task_tracker_back/pkg/config"
	"github.com/Task_tracker_back/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func NewDatabaseConnection(cfg *config.Config) *gorm.DB {
	db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			cfg.DB.Host,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.DBName,
			cfg.DB.Port,
			cfg.DB.SSLMode,
			cfg.DB.Timezone)),
		&gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Database connection successful")

	db.AutoMigrate(&models.User{})

	return db
}
