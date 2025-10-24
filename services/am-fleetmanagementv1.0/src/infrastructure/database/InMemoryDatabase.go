package database

import (
	"fleetmanagement/infrastructure/database/entities"
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"

	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var InMemDB *gorm.DB

func ConnectDB(isInMemory bool) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if isInMemory {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {

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

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	// Connecting to SQLite in-memory database
	if err != nil {
		return nil, err
	}

	InMemDB = db
	return db, nil
}

func Migrate() error {
	err := InMemDB.AutoMigrate(&entities.FleetPersistenceEntity{})
	if err != nil {
		return err
	}

	err = InMemDB.AutoMigrate(&entities.CarPersistenceEntity{}) // Add your model structs here
	if err != nil {
		return err
	}

	return nil

}

func CloseDB() error {
	sqlDB, err := InMemDB.DB()
	if err == nil {
		err = sqlDB.Close()
	}
	if err != nil {
		msg := fmt.Sprintf("Couldn't close database")
		log.Error(msg, err)
		return fmt.Errorf("%s: %w", msg, err)
	}
	return nil
}
