package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationError holds validation error information
type ValidationError struct {
	Field   string
	Message string
	Tag     string
}

// ParameterChecker validates request parameters
type ParameterChecker struct {
	Validate *validator.Validate
}

// NewParameterChecker creates a new parameter checker
func NewParameterChecker() *ParameterChecker {
	v := validator.New()
	return &ParameterChecker{
		Validate: v,
	}
}

// ValidateStruct validates a struct using go-playground/validator
func (pc *ParameterChecker) ValidateStruct(s interface{}) []ValidationError {
	errs := pc.Validate.Struct(s)
	if errs == nil {
		return nil
	}

	var validationErrs []ValidationError
	for _, err := range errs.(validator.ValidationErrors) {
		validationErrs = append(validationErrs, ValidationError{
			Field:   err.Field(),
			Message: err.Error(),
			Tag:     err.Tag(),
		})
	}
	return validationErrs
}

// CheckRequired validates that required fields are present
func CheckRequired(params map[string]interface{}, required []string) []string {
	var missing []string
	for _, field := range required {
		if _, ok := params[field]; !ok || params[field] == "" {
			missing = append(missing, field)
		}
	}
	return missing
}

// ValidateParams validates parameters based on their type and rules
func ValidateParams(param interface{}) error {
	pc := NewParameterChecker()
	errs := pc.ValidateStruct(param)
	if len(errs) > 0 {
		var errMsgs []string
		for _, e := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", e.Field, e.Message))
		}
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, ", "))
	}
	return nil
}

// GetValidationErrors extracts validation errors from validator errors
func GetValidationErrors(errs validator.ValidationErrors) map[string]string {
	result := make(map[string]string)
	for _, err := range errs {
		result[err.Field()] = err.Tag()
	}
	return result
}

// IsRequired checks if a field is marked as required via binding tag
func IsRequired(field reflect.StructField) bool {
	tag := field.Tag.Get("binding")
	return strings.Contains(tag, "required")
}
