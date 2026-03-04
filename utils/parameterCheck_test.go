package utils

import (
	"testing"
)

type TestUser struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"gte=0,lte=150"`
}

func TestParameterChecker_ValidateStruct(t *testing.T) {
	checker := NewParameterChecker()
	
	// Test valid struct
	user := TestUser{Name: "John", Age: 25}
	errs := checker.ValidateStruct(user)
	if len(errs) > 0 {
		t.Errorf("Expected no errors for valid struct, got %d", len(errs))
	}
}

func TestCheckRequired(t *testing.T) {
	params := map[string]interface{}{"name": "John", "age": 25}
	required := []string{"name", "age"}
	
	missing := CheckRequired(params, required)
	if len(missing) != 0 {
		t.Errorf("Expected 0 missing, got %d", len(missing))
	}
	
	// Test missing
	missing = CheckRequired(params, []string{"email"})
	if len(missing) != 1 {
		t.Errorf("Expected 1 missing, got %d", len(missing))
	}
}
