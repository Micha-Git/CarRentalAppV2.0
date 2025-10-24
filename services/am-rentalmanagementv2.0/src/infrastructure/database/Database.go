package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"rentalmanagement/infrastructure/database/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		postgresHost = "postgres"
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		postgresPort = "5432"
	}

	postgresDatabase := os.Getenv("POSTGRES_DATABASE")
	if postgresDatabase == "" {
		postgresDatabase = "postgres"
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		postgresUser = "postgres"
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		postgresPassword = "postgres"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db

	return db, nil
}

func Migrate() error {
	err := DB.AutoMigrate(&entities.RentalPersistenceEntity{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&entities.RentableCarPersistenceEntity{})
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err == nil {
		err = sqlDB.Close()
	}

	if err != nil {
		msg := fmt.Sprintf("Couldn't close database")
		log.Error(msg, ": ", err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}
