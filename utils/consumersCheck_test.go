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

func TestConsumersCheck_Check(t *testing.T) {
	tests := []struct {
		name            string
		acceptHeader    string
		supportedTypes  []string
		expect          bool
	}{
		{
			name:           "Accept matches supported type",
			acceptHeader:   "application/json",
			supportedTypes: []string{"application/json"},
			expect:         true,
		},
		{
			name:           "Accept matches one of supported types",
			acceptHeader:   "application/xml",
			supportedTypes: []string{"application/json", "application/xml"},
			expect:         true,
		},
		{
			name:           "No Accept header accepts all",
			acceptHeader:   "",
			supportedTypes: []string{"application/json"},
			expect:         true,
		},
		{
			name:           "Accept does not match",
			acceptHeader:   "text/plain",
			supportedTypes: []string{"application/json"},
			expect:         false,
		},
		{
			name:           "Wildcard accept",
			acceptHeader:   "*/*",
			supportedTypes: []string{"application/json"},
			expect:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.acceptHeader != "" {
				req.Header.Set("Accept", tt.acceptHeader)
			}
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			checker := NewConsumersCheck(tt.supportedTypes)
			result := checker.Check(c)

			if result != tt.expect {
				t.Errorf("Expected %v, got %v", tt.expect, result)
			}
		})
	}
}

func TestConsumersCheck_GetAcceptedType(t *testing.T) {
	tests := []struct {
		name           string
		acceptHeader   string
		supportedTypes []string
		expect         string
	}{
		{
			name:           "Returns supported type",
			acceptHeader:   "application/json",
			supportedTypes: []string{"application/json", "application/xml"},
			expect:         "application/json",
		},
		{
			name:           "No Accept returns first supported",
			acceptHeader:   "",
			supportedTypes: []string{"application/json", "application/xml"},
			expect:         "application/json",
		},
		{
			name:           "No Accept and no supported returns default",
			acceptHeader:   "",
			supportedTypes: []string{},
			expect:         "*/*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.acceptHeader != "" {
				req.Header.Set("Accept", tt.acceptHeader)
			}
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			checker := NewConsumersCheck(tt.supportedTypes)
			result := checker.GetAcceptedType(c)

			if result != tt.expect {
				t.Errorf("Expected '%s', got '%s'", tt.expect, result)
			}
		})
	}
}

func TestCheckProduce(t *testing.T) {
	tests := []struct {
		name            string
		contentType     string
		supportedTypes  []string
		expect          bool
	}{
		{
			name:           "Content-Type matches",
			contentType:    "application/json",
			supportedTypes: []string{"application/json"},
			expect:         true,
		},
		{
			name:           "Content-Type does not match",
			contentType:    "text/plain",
			supportedTypes: []string{"application/json"},
			expect:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Set("Content-Type", tt.contentType)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			result := CheckProduce(c, tt.supportedTypes)
			if result != tt.expect {
				t.Errorf("Expected %v, got %v", tt.expect, result)
			}
		})
	}
}
