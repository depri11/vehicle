package database

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/depri11/vehicle/src/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	env := os.Getenv("APP_ENV")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed connecting to database")
	}

	dbLife, err := db.DB()
	if err != nil {
		return nil, errors.New("failed connecting to database")
	}

	if env == "development" {
		db.AutoMigrate(&models.User{}, &models.Historys{}, &models.Vehicle{}, &models.VehicleImage{})
	}

	dbLife.SetConnMaxIdleTime(10)
	dbLife.SetMaxOpenConns(100)
	dbLife.SetConnMaxLifetime(time.Hour)

	return db, nil
}
