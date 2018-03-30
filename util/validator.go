package util

import (
	"reflect"

	validator "gopkg.in/go-playground/validator.v8"
	gvalidator "gopkg.in/go-playground/validator.v9"
)

type defaultValidator struct {
	validate *gvalidator.Validate
}

var dValidator *defaultValidator = newValidator()

func Validator() defaultValidator {
	return *dValidator
}

func ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		if err := dValidator.validate.Struct(obj); err != nil {
			return error(err)
		}
	}
	return nil
}

func (v defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}
	return nil
}

func (v defaultValidator) RegisterValidation(string, validator.Func) error {
	return nil
}

func newValidator() *defaultValidator {
	valid := gvalidator.New()
	valid.SetTagName("validate")
	return &defaultValidator{
		validate: valid,
	}
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
