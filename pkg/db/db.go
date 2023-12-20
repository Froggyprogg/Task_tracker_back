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

// DisableForeignKeyConstraintWhenMigrating: true,
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
		&gorm.Config{DisableForeignKeyConstraintWhenMigrating: false})

	if err != nil {
		panic(err)
	}

	log.Println("Database connection successful")

	db.AutoMigrate(&models.User{}, &models.Board{}, &models.Role{}, &models.Column{}, &models.UserBoard{}, &models.Task{}, &models.Subtask{}, &models.Status{}, &models.Report{}, &models.Comment{})

	return db
}
