package main

import (
	"os"
	"strconv"

	"github.com/Luqky07/user_api/data"
	"github.com/Luqky07/user_api/models/tables"
	"github.com/Luqky07/user_api/routes"
	"github.com/Luqky07/user_api/services"
	"github.com/Luqky07/user_api/validators"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// Main function
func main() {
	//Load the variables of .env
	if varErr := godotenv.Load(); varErr != nil {
		//If it return an error throw panic
		panic(varErr)
	}

	data.API_ADMIN = os.Getenv("API_ADMIN_ID")
	if data.API_ADMIN == "" {
		panic("API_ADMIN variable is not defined")
	}

	services.MAIL = os.Getenv("MAIL")
	if services.MAIL == "" {
		panic("MAIL variable is not defined")
	}

	services.MAIL_PASSWORD = os.Getenv("MAIL_PASSWORD")
	if services.MAIL_PASSWORD == "" {
		panic("MAIL_PASSWORD variable is not defined")
	}

	services.SMTP = os.Getenv("SMTP")
	if services.SMTP == "" {
		panic("SMTP variable is not defined")
	}

	services.PORT, _ = strconv.Atoi(os.Getenv("MAIL_PORT"))
	if services.PORT == 0 {
		panic("MAIL_PORT variable is not defined")
	}

	//Set database connection
	data.ConnectDatabase()

	//Migrating the database models
	err := data.DB.AutoMigrate(
		&tables.EmployeeRole{},
		&tables.Employee{},
		&tables.PhoneType{},
		&tables.Phone{},
		&tables.EmployeeCredential{},
		&tables.EmployeePhone{},
		&tables.Client{},
		&tables.ClientPhone{},
		&tables.ClientEmailVerification{},
		&tables.ClientCredential{},
	)

	//If the migration return an error throw panic
	if err != nil {
		panic(err)
	}

	//Create the router
	router := gin.Default()

	//Set the custom validations
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validators.RegisterCreateUserValidations(v)
	}

	//Setup the api routes
	routes.SetupRoutes(router)

	//Setup the api
	router.Run("localhost:8080")
}
