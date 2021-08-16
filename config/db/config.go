package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=golightweb port=5432"
	var err error
	connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Connection() *gorm.DB {
	return connection
}
