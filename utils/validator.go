package utils

import (
	"fmt"
	"reflect"
)

const tagName = "validate"

type Validator interface {
	Validate(interface{}) error
}

/**/
type NilValidator struct {
}

func (validatore *NilValidator) Validate(f interface{}) error {
	if reflect.TypeOf(f) == nil {
		return fmt.Errorf("is nil")
	}
	return nil
}

func ValidateStruct(s interface{}) []error {
	errs := []error{}

	// value representing concrete value of the interface
	value := reflect.ValueOf(s)

	for i := 0; i < value.NumField(); i++ {
		// retrieving the tag from field
		// the field is retrived as value of Type()
		tag := value.Type().Field(i).Tag.Get(tagName)

		// check for empty tag, empty continue the iteration
		if tag == "" {
			continue
		}
		// get the validator to perform type validation
		validator := getValidator(tag)

		valueInterface := value.Field(i).Interface()
		err := validator.Validate(valueInterface)

		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %s", value.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}

func getValidator(tag string) Validator {
	switch tag {
	case "NotNil":
		return &NilValidator{}
	default:
		return &NilValidator{}
	}
}
