package main

import (
	"fmt"

	"github.com/Luqky07/user_api/data"
	"github.com/Luqky07/user_api/models/tables"
	"github.com/Luqky07/user_api/routes"
	"github.com/Luqky07/user_api/validators"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	if varErr := godotenv.Load(); varErr != nil {
		panic(varErr)
	}

	db, err := data.Database()
	if err != nil {
		return
	}

	err = db.AutoMigrate(
		&tables.EmployeeRole{},
		&tables.Employee{},
		&tables.Client{},
	)
	if err != nil {
		fmt.Println("Error al migrar el esquema:", err)
		return
	}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validators.RegisterCreateUserValidations(v)
	}

	routes.SetupRoutes(router)

	router.Run("localhost:8080")
}
