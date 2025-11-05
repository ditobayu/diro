package database

import (
	"fmt"
	"log"

	mysqlgorm "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"diro-be/internal/config"
	"diro-be/internal/models"
)

// DB is the global database connection
var DB *gorm.DB

// Connect establishes a connection to the MySQL database
func Connect(cfg *config.Config) error {
	dsn := cfg.GetDBDSN()
	db, err := gorm.Open(mysqlgorm.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	log.Println("Connected to database successfully")

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Court{}, &models.Timeslot{}, &models.Reservation{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
