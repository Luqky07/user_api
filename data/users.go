package data

import (
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database = func() (*gorm.DB, error) {
	dsn := os.Getenv("USERS_DB")
	if dsn == "" {
		return nil, errors.New("la variable de entorno USERS_DB no está configurada")
	}
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		return db, err
	}
}
