package validators

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validate a spanish personal document
func validDocumentId(f1 validator.FieldLevel) bool {

	//Letters to validate the personal document
	letrasDni := [23]string{
		"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B",
		"N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E",
	}

	//Letters to validate the foreign document
	letrasNie := map[rune]rune{
		'X': '0',
		'Y': '1',
		'Z': '2',
	}

	//Convert the field to string using uppercase
	documentId := strings.ToUpper(f1.Field().String())

	//Regular expresion to validate a spanish personal document
	regexDocumentId := regexp.MustCompile(`^(\d{8}[A-Z]|[XYZ]\d{7}[A-Z])$`)

	//If the regular expresion do not validate the field return false
	if !regexDocumentId.MatchString(documentId) {
		return false
	}

	//Create a slice using the field letters
	documentSlice := []rune(documentId)

	//Convert de foreign document letters to the corresponding number
	for key, value := range letrasNie {
		if documentSlice[0] == key {
			documentSlice[0] = value
			break
		}
	}

	//Save the personal document letter
	letter := string(documentSlice[len(documentSlice)-1])

	//Save the numbers of the personal document
	numbers, err := strconv.Atoi(string(documentSlice[:len(documentSlice)-1]))

	//If the numbers conversion make an error return false
	if err != nil {
		return false
	}

	//Save de control number of the personal document
	control := numbers % 23

	//Return the comparation of the letter of the document and the true letter
	return letter == letrasDni[control]
}

// Validate a spanish phone number
func validPhone(f1 validator.FieldLevel) bool {
	//Regular expresion of a spanish phone
	regexPhone := regexp.MustCompile(`^[6789]\d{8}$`)

	//Convert de field to a string
	phone := f1.Field().String()

	//Retun the regular expresion validation of the field
	return regexPhone.MatchString(phone)
}

// Registering the validating fuctions
func RegisterCreateUserValidations(v *validator.Validate) {
	v.RegisterValidation("validDocument", validDocumentId)
	v.RegisterValidation("validPhone", validPhone)
}
