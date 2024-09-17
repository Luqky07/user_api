package data

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var API_ADMIN string

// Connection with users database using dotenv
func ConnectDatabase() {

	//Get environment variable to connect the database
	dsn := os.Getenv("USERS_DB")

	//If the variable is empty return an error
	if dsn == "" {
		panic("USERS_DB variable is not defined")
	}

	//Set users database connection
	var err error
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		//If it fail return an error
		panic(err)
	}
}
