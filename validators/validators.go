package validators

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validDocumentId(f1 validator.FieldLevel) bool {
	letrasDni := [23]string{
		"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B",
		"N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E",
	}

	letrasNie := map[rune]rune{
		'X': '0',
		'Y': '1',
		'Z': '2',
	}

	documentId := strings.ToUpper(f1.Field().String())

	regexDocumentId := regexp.MustCompile(`^(\d{8}[A-Z]|[XYZ]\d{7}[A-Z])$`)

	if !regexDocumentId.MatchString(documentId) {
		return false
	}

	documentSlice := []rune(documentId)

	for key, value := range letrasNie {
		if documentSlice[0] == key {
			documentSlice[0] = value
			break
		}
	}

	letra := string(documentSlice[len(documentSlice)-1])

	numeros, err := strconv.Atoi(string(documentSlice[:len(documentSlice)-1]))

	fmt.Println(numeros)

	if err != nil {
		return false
	}

	control := numeros % 23

	return letra == letrasDni[control]
}

func validPhone(f1 validator.FieldLevel) bool {
	regexPhone := regexp.MustCompile(`^[6789]\d{8}$`)

	phone := f1.Field().String()

	return regexPhone.MatchString(phone)
}

func RegisterCreateUserValidations(v *validator.Validate) {
	v.RegisterValidation("validDocument", validDocumentId)
	v.RegisterValidation("validPhone", validPhone)
}
