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
	// Create test request with JSON body
	body := `{"name": "test", "age": 25}`
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create test struct
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
	
	if user.Age != 25 {
		t.Errorf("Expected age 25, got %d", user.Age)
	}
}

func TestBindParamsFromSource(t *testing.T) {
	tests := []struct {
		name         string
		source       ParameterBindingSource
		query        string
		header       string
		body         string
		contentType  string
		expectErr    bool
		checkField   func(*testing.T, *TestUser)
	}{
		{
			name:        "Bind from query",
			source:      BindingFromQuery,
			query:       "name=querytest&age=30",
			expectErr:   false,
			checkField: func(t *testing.T, u *TestUser) {
				if u.Name != "querytest" {
					t.Errorf("Expected name 'querytest', got '%s'", u.Name)
				}
			},
		},
	}

	type TestUser struct {
		Name string `json:"name" form:"name" query:"name"`
		Age  int    `json:"age" form:"age" query:"age"`
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test?"+tt.query, nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			user := &TestUser{}
			err := BindParamsFromSource(c, tt.source, user)
			
			if tt.expectErr && err == nil {
				t.Error("Expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if tt.checkField != nil {
				tt.checkField(t, user)
			}
		})
	}
}

func TestBindStruct(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		contentType string
		expectName  string
		expectAge   int
	}{
		{
			name:        "JSON body",
			body:        `{"name": "jsontest", "age": 20}`,
			contentType: "application/json",
			expectName:  "jsontest",
			expectAge:   20,
		},
		{
			name:        "Form data",
			body:        "name=formtest&age=35",
			contentType: "application/x-www-form-urlencoded",
			expectName:  "formtest",
			expectAge:   35,
		},
	}

	type TestUser struct {
		Name string `json:"name" form:"name"`
		Age  int    `json:"age" form:"age"`
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", tt.contentType)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			user := &TestUser{}
			err := BindStruct(c, user)
			
			if err != nil {
				t.Errorf("BindStruct failed: %v", err)
			}
			
			if user.Name != tt.expectName {
				t.Errorf("Expected name '%s', got '%s'", tt.expectName, user.Name)
			}
		})
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
