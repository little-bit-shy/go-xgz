package validator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// V validator data
func V(data interface{}) (errTag interface{}, errString interface{}) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Name
	})
	if errs := validate.Struct(data); errs != nil {
		filedError := errs.(validator.ValidationErrors)[0].(validator.FieldError)
		errTag = filedError.StructField()
		errString = filedError.Namespace()
	}
	return
}
