package repository

import (
	"fmt"
	"github.com/oguzhand95/go-web-service-starter/src/internal/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
}

func NewDatabase(databaseConfiguration *configuration.DatabaseConfiguration) (*Database, error) {
	db := &Database{}
	var err error

	db.Gorm, err = db.runPostgresql(databaseConfiguration)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Database) Close() error {
	sqlDB, err := d.Gorm.DB()

	if err != nil {
		return err
	}

	err = sqlDB.Close()

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) runPostgresql(databaseConfiguration *configuration.DatabaseConfiguration) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		databaseConfiguration.Host, databaseConfiguration.Username, databaseConfiguration.Password,
		databaseConfiguration.DatabaseName, databaseConfiguration.Port)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
