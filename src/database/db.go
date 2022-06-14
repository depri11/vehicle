package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/depri11/vehicle/src/modules/v1/history"
	"github.com/depri11/vehicle/src/modules/v1/users"
	vehicle "github.com/depri11/vehicle/src/modules/v1/vehicles"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	env := os.Getenv("APP_ENV")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed connecting to database")
	}

	dbLife, err := db.DB()
	if err != nil {
		return nil, errors.New("failed connecting to database")
	}

	if env == "development" {
		db.AutoMigrate(&users.User{}, &history.Historys{}, &vehicle.Vehicle{}, &vehicle.VehicleImage{})
	}

	dbLife.SetConnMaxIdleTime(10)
	dbLife.SetMaxOpenConns(100)
	dbLife.SetConnMaxLifetime(time.Hour)

	return db, nil
}
