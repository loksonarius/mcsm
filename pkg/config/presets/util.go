package presets

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("divby", divisibleBy)
	validate.RegisterValidation("port", isPort)
}

func divisibleBy(fl validator.FieldLevel) bool {
	switch fl.Field().Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		p, err := strconv.ParseInt(fl.Param(), 10, 64)
		if err != nil {
			return false
		}

		i := fl.Field().Int()
		return i%p == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		p, err := strconv.ParseUint(fl.Param(), 10, 64)
		if err != nil {
			return false
		}

		i := fl.Field().Uint()
		return i%p == 0
	default:
		panic(fmt.Sprintf("Bad field type %T", fl.Field().Interface()))
	}
}

func isPort(fl validator.FieldLevel) bool {
	switch fl.Field().Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		i := fl.Field().Int()
		return i >= 1 && i <= 65535
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		i := fl.Field().Uint()
		return i >= 1 && i <= 65535
	default:
		panic(fmt.Sprintf("Bad field type %T", fl.Field().Interface()))
	}
}
