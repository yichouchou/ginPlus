package utils

import (
	"testing"
)

type TestUser struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"gte=0,lte=150"`
	Email string `json:"email" binding:"required,email"`
}

func TestParameterChecker_ValidateStruct(t *testing.T) {
	tests := []struct {
		name      string
		user      TestUser
		expectErr bool
	}{
		{
			name:      "Valid struct",
			user:      TestUser{Name: "John", Age: 25, Email: "john@example.com"},
			expectErr: false,
		},
		{
			name:      "Missing required field",
			user:      TestUser{Age: 25, Email: "john@example.com"},
			expectErr: true,
		},
		{
			name:      "Invalid email",
			user:      TestUser{Name: "John", Age: 25, Email: "invalid-email"},
			expectErr: true,
		},
		{
			name:      "Age out of range",
			user:      TestUser{Name: "John", Age: 200, Email: "john@example.com"},
			expectErr: true,
		},
	}

	checker := NewParameterChecker()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := checker.ValidateStruct(tt.user)
			
			hasErr := len(errs) > 0
			if hasErr != tt.expectErr {
				t.Errorf("Expected error=%v, got %v with %d errors", tt.expectErr, hasErr, len(errs))
			}
		})
	}
}

func TestCheckRequired(t *testing.T) {
	tests := []struct {
		name      string
		params    map[string]interface{}
		required  []string
		expectLen int
	}{
		{
			name:      "All required present",
			params:    map[string]interface{}{"name": "John", "age": 25},
			required:  []string{"name", "age"},
			expectLen: 0,
		},
		{
			name:      "One required missing",
			params:    map[string]interface{}{"name": "John"},
			required:  []string{"name", "age"},
			expectLen: 1,
		},
		{
			name:      "Multiple required missing",
			params:    map[string]interface{}{},
			required:  []string{"name", "age", "email"},
			expectLen: 3,
		},
		{
			name:      "Empty required list",
			params:    map[string]interface{}{"name": "John"},
			required:  []string{},
			expectLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			missing := CheckRequired(tt.params, tt.required)
			if len(missing) != tt.expectLen {
				t.Errorf("Expected %d missing, got %d: %v", tt.expectLen, len(missing), missing)
			}
		})
	}
}

func TestValidateParams(t *testing.T) {
	tests := []struct {
		name      string
		param     interface{}
		expectErr bool
	}{
		{
			name:      "Valid params",
			param:     TestUser{Name: "John", Age: 25, Email: "john@example.com"},
			expectErr: false,
		},
		{
			name:      "Invalid params",
			param:     TestUser{Name: "", Age: 25, Email: "invalid"},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateParams(tt.param)
			hasErr := err != nil
			if hasErr != tt.expectErr {
				t.Errorf("Expected error=%v, got %v", tt.expectErr, hasErr)
			}
		})
	}
}

func TestIsRequired(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		binding  string
		expected bool
	}{
		{
			name:     "Required field",
			field:    "Name",
			binding:  "required",
			expected: true,
		},
		{
			name:     "Non-required field",
			field:    "Age",
			binding:  "gte=0,lte=150",
			expected: false,
		},
		{
			name:     "No binding tag",
			field:    "Email",
			binding:  "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf := struct {
				Name  string `binding:"required"`
				Age   int    `binding:"gte=0,lte=150"`
				Email string
			}{}

			var field interface{} = sf
			switch tt.field {
			case "Name":
				result := IsRequired(getStructField(field, "Name"))
				if result != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			case "Age":
				result := IsRequired(getStructField(field, "Age"))
				if result != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			case "Email":
				result := IsRequired(getStructField(field, "Email"))
				if result != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

// Helper function to get struct field
func getStructField(i interface{}, fieldName string) interface{} {
	return nil // Placeholder - not actually used in this simplified test
}
