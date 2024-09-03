package database

import (
	"fmt"
	"log"

	"github.com/Frhnmj2004/restaurant-admin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.FoodItem{},
		&models.Grocery{},
		&models.Ingredient{},
		&models.Order{},
		&models.Revenue{},
		&models.User{},
	)

	if err != nil {
		log.Printf("Migration failed: %v", err)
		return err
	}

	log.Println("Migration successful")
	return nil
}
