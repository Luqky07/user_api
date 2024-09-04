package controllers

import (
	"fmt"
	"net/http"

	"github.com/Luqky07/user_api/models/request"
	"github.com/Luqky07/user_api/models/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}

func CreateNewUser(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// Verificamos si es un error de validación
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErrs {
				// Personalizar mensajes por tag
				switch fieldError.Tag() {
				case "required":
					errorMessages = append(errorMessages, fmt.Sprintf("El campo %s es obligatorio.", fieldError.Field()))
				case "validDocument":
					errorMessages = append(errorMessages, "El documento de identidad no es correcto")
				case "email":
					errorMessages = append(errorMessages, fmt.Sprintf("El campo %s debe ser un correo electrónico válido.", fieldError.Field()))
				case "validPhone":
					errorMessages = append(errorMessages, "El teléfono no tiene un formato correcto")
				default:
					errorMessages = append(errorMessages, fmt.Sprintf("El campo %s tiene un error de validación.", fieldError.Field()))
				}
			}
			// Devolver mensajes de error personalizados
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		// Si es otro tipo de error, lo devolvemos tal cual
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := response.NewCreateUserResponse(req)

	c.JSON(http.StatusCreated, res)
}
