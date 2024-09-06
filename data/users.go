package data

import (
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection with users database using dotenv
var Database = func() (*gorm.DB, error) {

	//Get environment variable to connect the database
	dsn := os.Getenv("USERS_DB")

	//If the variable is empty return an error
	if dsn == "" {
		return nil, errors.New("USERS_DB variable is not defined")
	}

	//Set users database connection
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		//If it fail return an error
		return nil, err
	} else {
		//Return the database
		return db, err
	}
}
