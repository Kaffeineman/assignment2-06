package database

import (
	"fmt"

	"github.com/Kaffeineman/assignment2-06/assignment2/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func InitDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(model.Item{}, model.Order{})

	return db, nil
}
