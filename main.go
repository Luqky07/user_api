package main

import (
	"github.com/Luqky07/user_api/data"
	"github.com/Luqky07/user_api/models/tables"
	"github.com/Luqky07/user_api/routes"
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

	//Set database connection
	db, err := data.Database()
	if err != nil {
		//If it return an error throw panic
		panic(err)
	}

	//Migrating the database models
	err = db.AutoMigrate(
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
