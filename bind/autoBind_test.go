package bind

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestAutoBindParams(t *testing.T) {
	body := `{"name": "test", "age": 25}`
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	type TestUser struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	user := &TestUser{}
	err := AutoBindParams(c, user)
	
	if err != nil {
		t.Errorf("AutoBindParams failed: %v", err)
	}
	
	if user.Name != "test" {
		t.Errorf("Expected name 'test', got '%s'", user.Name)
	}
}

func TestAutoBind_GET(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test?name=autotest&age=40", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	type TestUser struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	user := &TestUser{}
	err := AutoBind(c, user)
	
	if err != nil {
		t.Errorf("AutoBind GET failed: %v", err)
	}
	
	if user.Name != "autotest" {
		t.Errorf("Expected name 'autotest', got '%s'", user.Name)
	}
}

func TestAutoBind_POST_JSON(t *testing.T) {
	body := `{"name": "posttest", "age": 45}`
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	type TestUser struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &TestUser{}
	err := AutoBind(c, user)
	
	if err != nil {
		t.Errorf("AutoBind POST JSON failed: %v", err)
	}
	
	if user.Name != "posttest" {
		t.Errorf("Expected name 'posttest', got '%s'", user.Name)
	}
}
