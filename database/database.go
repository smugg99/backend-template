package database

import (
	"fmt"

	"github.com/B-P-Coders/backend-template/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw",
		config.Conf.DbConfig.Host,
		config.Conf.DbConfig.User,
		config.Conf.DbConfig.Password,
		"kujonhouse",
		"5432",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate()

	return db, nil
}
