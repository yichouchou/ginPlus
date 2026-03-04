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

func TestProducesSet_SetContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ps := NewProducesSet([]string{"application/json", "application/xml"})
	ps.SetContentType(c)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected 'application/json', got '%s'", contentType)
	}
}

func TestProducesSet_GetContentType(t *testing.T) {
	tests := []struct {
		name           string
		acceptHeader   string
		contentTypes   []string
		expect         string
	}{
		{
			name:         "Accept matches content type",
			acceptHeader: "application/json",
			contentTypes: []string{"application/json", "application/xml"},
			expect:       "application/json",
		},
		{
			name:         "No Accept returns first content type",
			acceptHeader: "",
			contentTypes: []string{"application/json", "application/xml"},
			expect:       "application/json",
		},
		{
			name:         "Wildcard accept",
			acceptHeader: "*/*",
			contentTypes: []string{"application/json"},
			expect:       "application/json",
		},
		{
			name:         "No content types defaults to JSON",
			acceptHeader: "",
			contentTypes: []string{},
			expect:       "application/json",
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

			ps := NewProducesSet(tt.contentTypes)
			result := ps.GetContentType(c)

			if result != tt.expect {
				t.Errorf("Expected '%s', got '%s'", tt.expect, result)
			}
		})
	}
}

func TestContentTypeConstants(t *testing.T) {
	if ContentTypeJSON != "application/json" {
		t.Errorf("Expected 'application/json', got '%s'", ContentTypeJSON)
	}
	if ContentTypeXML != "application/xml" {
		t.Errorf("Expected 'application/xml', got '%s'", ContentTypeXML)
	}
	if ContentTypeHTML != "text/html" {
		t.Errorf("Expected 'text/html', got '%s'", ContentTypeHTML)
	}
	if ContentTypeText != "text/plain" {
		t.Errorf("Expected 'text/plain', got '%s'", ContentTypeText)
	}
	if ContentTypeForm != "application/x-www-form-urlencoded" {
		t.Errorf("Expected 'application/x-www-form-urlencoded', got '%s'", ContentTypeForm)
	}
	if ContentTypeMulti != "multipart/form-data" {
		t.Errorf("Expected 'multipart/form-data', got '%s'", ContentTypeMulti)
	}
}
