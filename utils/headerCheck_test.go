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
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	validator := NewHeaderValidator([]string{"Content-Type", "Accept"}, []string{})
	result := validator.ValidateHeaders(c)

	if !result.Valid {
		t.Error("Expected valid result")
	}
}

func TestCheckContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	result := CheckContentType(c, []string{"application/json"})
	if !result {
		t.Error("Expected true for matching content type")
	}
}

func TestHasHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Custom-Header", "exists")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	if !HasHeader(c, "X-Custom-Header") {
		t.Error("Expected HasHeader to return true")
	}
}
