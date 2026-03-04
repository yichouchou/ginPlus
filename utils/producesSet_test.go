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

	ps := NewProducesSet([]string{"application/json"})
	ps.SetContentType(c)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected 'application/json', got '%s'", contentType)
	}
}

func TestContentTypeConstants(t *testing.T) {
	if ContentTypeJSON != "application/json" {
		t.Errorf("Expected 'application/json', got '%s'", ContentTypeJSON)
	}
}
