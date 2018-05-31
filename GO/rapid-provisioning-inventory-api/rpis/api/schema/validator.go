package schema

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidationError is the common error container
type ValidationError struct {
	FieldName        string `json:"field"`
	ErrorDescription string `json:"description"`
}

func (err ValidationError) Error() string {
	return fmt.Sprintf("'%s' %s", err.FieldName, err.ErrorDescription)
}

// Validator is the field type validator
type Validator interface {
	Validate(interface{}, string) error
}

type numberValidator struct {
	Min int
	Max int
}

func (n numberValidator) Validate(val interface{}, name string) error {
	v := val.(int)
	if v < n.Min {
		return ValidationError{FieldName: name, ErrorDescription: "is required"}
	}
	return nil
}

type stringValidator struct {
	Min int
	Max int
}

func (s stringValidator) Validate(val interface{}, name string) error {
	v := val.(string)
	if len(v) < s.Min {
		return ValidationError{FieldName: name, ErrorDescription: "is required"}
	}
	if s.Max > 0 && len(v) > s.Max {
		return ValidationError{FieldName: name, ErrorDescription: "exceeds max length"}
	}
	return nil
}

type enumValidator struct {
	Min        int
	EnumValues string
	enumList   []string
}

func (e enumValidator) Validate(val interface{}, name string) error {
	v := val.(string)
	if len(v) < e.Min {
		return ValidationError{FieldName: name, ErrorDescription: "is required"}
	}
	e.enumList = strings.Split(e.EnumValues, ",")
	for _, x := range e.enumList {
		if x == v {
			return nil
		}
	}
	return ValidationError{FieldName: name, ErrorDescription: fmt.Sprintf("value should be one of %s", e.EnumValues)}
}

type defaultValidator struct{}

func (d defaultValidator) Validate(_ interface{}, _ string) error {
	return nil
}

func getValidator(f reflect.StructField) Validator {
	tag := strings.SplitN(f.Tag.Get("validate"), ",", 2)
	if len(tag) > 1 {
		switch tag[0] {
		case "number":
			v := numberValidator{}
			fmt.Sscanf(tag[1], "min=%d,max=%d", &v.Min, &v.Max)
			return v
		case "string":
			v := stringValidator{}
			fmt.Sscanf(tag[1], "min=%d,max=%d", &v.Min, &v.Max)
			return v
		case "enum":
			v := enumValidator{}
			fmt.Sscanf(tag[1], "min=%d,%v", &v.Min, &v.EnumValues)
			return v
		}
	}
	return defaultValidator{}
}

//ValidateFields is a generic field validator based on tags
func ValidateFields(parent string, ins interface{}) []error {
	allErrors := make([]error, 0)
	v := reflect.ValueOf(ins)
	// recover if reflect panics
	defer recover()
here:
	for i := 0; i < v.NumField(); i++ {
		tag := strings.SplitN(v.Type().Field(i).Tag.Get("validate"), ",", 2)
		if len(tag) > 0 && tag[0] == "-" {
			continue here
		}
		if v.Type().Field(i).Type.Kind().String() == "struct" {
			errs := ValidateFields(v.Type().Field(i).Name, v.Field(i).Interface())
			if len(errs) > 0 {
				allErrors = append(allErrors, errs...)
			}
		} else {
			validator := getValidator(v.Type().Field(i))
			f := v.Type().Field(i)
			fName := f.Tag.Get("json")
			if fName == "" {
				fName = v.Type().Field(i).Name
			}
			err := validator.Validate(v.Field(i).Interface(), parent+"."+fName)
			if err != nil {
				allErrors = append(allErrors, err)
			}
		}
	}
	return allErrors
}
