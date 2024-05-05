package utils

import (
	"log"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

func Hashpassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	return string(bytes), err
}

func ValidateStruct(s interface {}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		log.Printf("there was an error validating the user struct: %v", err)
		return err 
	}
	return nil
}

func SanitizeStructs(s interface{}) interface{} {
	p := bluemonday.UGCPolicy()
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Pointer{
		panic("SanitizeStructs only works on pointers")
	}
	val = val.Elem()					
	if val.Kind()!= reflect.Struct {
		panic("SanitizeStructs only works on structs")
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			sanitized := p.Sanitize(field.String())
			val.Field(i).SetString(sanitized)
		}
	}
	return s
}
