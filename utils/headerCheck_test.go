package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestHeaderValidator_ValidateHeaders(t *testing.T) {
	tests := []struct {
		name          string
		required      []string
		optional      []string
		headerSetup  func(*http.Request)
		expectValid   bool
		expectMessage string
	}{
		{
			name:     "All required headers present",
			required: []string{"Content-Type", "Accept"},
			optional: []string{"Authorization"},
			headerSetup: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Accept", "application/json")
				req.Header.Set("Authorization", "Bearer token123")
			},
			expectValid: true,
		},
		{
			name:     "Missing required header",
			required: []string{"Content-Type", "Accept"},
			optional: []string{"Authorization"},
			headerSetup: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
			},
			expectValid:   false,
			expectMessage: "Missing required header: Accept",
		},
		{
			name:     "No required headers",
			required: []string{},
			optional: []string{"Authorization"},
			headerSetup: func(req *http.Request) {
				req.Header.Set("Authorization", "Bearer token123")
			},
			expectValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.headerSetup != nil {
				tt.headerSetup(req)
			}
			
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			validator := NewHeaderValidator(tt.required, tt.optional)
			result := validator.ValidateHeaders(c)

			if result.Valid != tt.expectValid {
				t.Errorf("Expected valid=%v, got %v", tt.expectValid, result.Valid)
			}

			if tt.expectMessage != "" && result.Message != tt.expectMessage {
				t.Errorf("Expected message '%s', got '%s'", tt.expectMessage, result.Message)
			}
		})
	}
}

func TestCheckContentType(t *testing.T) {
	tests := []struct {
		name         string
		contentType  string
		allowedTypes []string
		expect       bool
	}{
		{
			name:         "JSON allowed",
			contentType:  "application/json",
			allowedTypes: []string{"application/json"},
			expect:       true,
		},
		{
			name:         "Multiple types allowed",
			contentType:  "application/xml",
			allowedTypes: []string{"application/json", "application/xml"},
			expect:       true,
		},
		{
			name:         "Type not allowed",
			contentType:  "text/plain",
			allowedTypes: []string{"application/json"},
			expect:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Set("Content-Type", tt.contentType)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			result := CheckContentType(c, tt.allowedTypes)
			if result != tt.expect {
				t.Errorf("Expected %v, got %v", tt.expect, result)
			}
		})
	}
}

func TestGetHeaderMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	headers := GetHeaderMap(c)

	if len(headers) != 3 {
		t.Errorf("Expected 3 headers, got %d", len(headers))
	}

	if headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type header, got '%s'", headers["Content-Type"])
	}

	if headers["Authorization"] != "Bearer token" {
		t.Errorf("Expected Authorization header, got '%s'", headers["Authorization"])
	}
}

func TestHasHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Custom-Header", "exists")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	if !HasHeader(c, "X-Custom-Header") {
		t.Error("Expected HasHeader to return true for existing header")
	}

	if HasHeader(c, "X-Non-Existent") {
		t.Error("Expected HasHeader to return false for non-existent header")
	}
}

func TestGetHeaderValue(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Test-Header", "test-value")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Test with existing header
	value := GetHeaderValue(c, "X-Test-Header", "default")
	if value != "test-value" {
		t.Errorf("Expected 'test-value', got '%s'", value)
	}

	// Test with non-existent header (fallback)
	value = GetHeaderValue(c, "X-Non-Existent", "fallback")
	if value != "fallback" {
		t.Errorf("Expected 'fallback', got '%s'", value)
	}
}
