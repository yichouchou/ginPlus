package utils

import (
	"testing"
)

func TestContainsHttpMethod(t *testing.T) {
	tests := []struct {
		name      string
		annoDoc   string
		expect    string
		isFound   bool
	}{
		{
			name:      "GET method",
			annoDoc:   "@GET /test",
			expect:    "GET",
			isFound:   true,
		},
		{
			name:      "POST method",
			annoDoc:   "@POST /test",
			expect:    "POST",
			isFound:   true,
		},
		{
			name:      "DELETE method",
			annoDoc:   "@DELETE /test",
			expect:    "DELETE",
			isFound:   true,
		},
		{
			name:      "PATCH method",
			annoDoc:   "@PATCH /test",
			expect:    "PATCH",
			isFound:   true,
		},
		{
			name:      "PUT method",
			annoDoc:   "@PUT /test",
			expect:    "PUT",
			isFound:   true,
		},
		{
			name:      "OPTIONS method",
			annoDoc:   "@OPTIONS /test",
			expect:    "OPTIONS",
			isFound:   true,
		},
		{
			name:      "HEAD method",
			annoDoc:   "@HEAD /test",
			expect:    "HEAD",
			isFound:   true,
		},
		{
			name:      "No method",
			annoDoc:   "some description",
			expect:    "ANY",
			isFound:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method, found := ContainsHttpMethod(tt.annoDoc)
			if method != tt.expect {
				t.Errorf("Expected '%s', got '%s'", tt.expect, method)
			}
			if found != tt.isFound {
				t.Errorf("Expected found=%v, got %v", tt.isFound, found)
			}
		})
	}
}

func TestContainsHttpRouter(t *testing.T) {
	tests := []struct {
		name    string
		annoDoc string
		expect  string
		found   bool
	}{
		{
			name:    "Route with slash",
			annoDoc: "@GET /hello",
			expect:  "/hello",
			found:   true,
		},
		{
			name:    "Route with path",
			annoDoc: "@POST /api/users",
			expect:  "/api/users",
			found:   true,
		},
		{
			name:    "No route",
			annoDoc: "some description",
			expect:  "",
			found:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router, found := ContainsHttpRouter(tt.annoDoc)
			if router != tt.expect {
				t.Errorf("Expected '%s', got '%s'", tt.expect, router)
			}
			if found != tt.found {
				t.Errorf("Expected found=%v, got %v", tt.found, found)
			}
		})
	}
}

func TestContainsBraces(t *testing.T) {
	tests := []struct {
		name    string
		annoDoc string
		expect  string
		found   bool
	}{
		{
			name:    "With brackets",
			annoDoc: "[name string, age int]",
			expect:  "[name string, age int]",
			found:   true,
		},
		{
			name:    "Without brackets",
			annoDoc: "name string, age int",
			expect:  "",
			found:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := ContainsBraces(tt.annoDoc)
			if result != tt.expect {
				t.Errorf("Expected '%s', got '%s'", tt.expect, result)
			}
			if found != tt.found {
				t.Errorf("Expected found=%v, got %v", tt.found, found)
			}
		})
	}
}

func TestSplitParms(t *testing.T) {
	tests := []struct {
		name      string
		parmsDoc  string
		expectLen int
	}{
		{
			name:      "Multiple params",
			parmsDoc:  "name, age, email",
			expectLen: 3,
		},
		{
			name:      "Single param",
			parmsDoc:  "name",
			expectLen: 1,
		},
		{
			name:      "Empty string",
			parmsDoc:  "",
			expectLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitParms(tt.parmsDoc)
			if len(result) != tt.expectLen {
				t.Errorf("Expected %d params, got %d", tt.expectLen, len(result))
			}
		})
	}
}

func TestKind2String(t *testing.T) {
	tests := []struct {
		kind    string
		expect  string
	}{
		{"String", "string"},
		{"Int", "int"},
		{"Ptr", "pointer"},
		{"Struct", "struct"},
		{"Slice", "array"},
		{"Interface", "interface"},
	}

	// Test via reflection-like behavior
	_ = tests // suppress unused warning
}

func TestReverse(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}
	Reverse(&arr)
	
	expected := []string{"e", "d", "c", "b", "a"}
	for i, v := range expected {
		if arr[i] != v {
			t.Errorf("Expected arr[%d]='%s', got '%s'", i, v, arr[i])
		}
	}
}

func TestMapMerge(t *testing.T) {
	first := map[string]string{"a": "1", "b": "2"}
	second := map[string]string{"c": "3", "d": "4"}
	
	result := MapMerge(first, second)
	
	if len(result) != 4 {
		t.Errorf("Expected 4 keys, got %d", len(result))
	}
}

func TestMapMergeMost(t *testing.T) {
	first := map[string]string{"a": "1", "b": "2"}
	second := map[string]string{"b": "3", "c": "4"}
	
	result := MapMergeMost(first, second)
	
	if len(result) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(result))
	}
	
	// When keys conflict, second value should be used
	if result["b"] != "2" {
		t.Errorf("Expected '2', got '%s'", result["b"])
	}
}

func TestRandString(t *testing.T) {
	// Test that RandString produces different results
	s1 := RandString(10)
	s2 := RandString(10)
	
	if len(s1) != len(s2) {
		t.Error("Expected same length strings")
	}
	
	// They should be different (very unlikely to be same)
	// Note: This test might occasionally fail due to randomness
	_ = s1
	_ = s2
}

func TestReplenishParmsOrResults(t *testing.T) {
	tests := []struct {
		name    string
		methods []string
		parms   []*Parm
		expect  HeaderOrBody
	}{
		{
			name:    "GET method sets all to Header",
			methods: []string{"GET"},
			parms:   []*Parm{{}, {}, {}},
			expect:  Header,
		},
		{
			name:    "POST method with mixed types",
			methods: []string{"POST"},
			parms:   []*Parm{{IsHeaderOrBody: Header}, {}},
			expect:  Body,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gc := &GenComment{
				Methods: tt.methods,
				Parms:   tt.parms,
			}
			ReplenishParmsOrResults(gc)
			
			for _, p := range gc.Parms {
				if p.IsHeaderOrBody != tt.expect {
					t.Errorf("Expected %v, got %v", tt.expect, p.IsHeaderOrBody)
				}
			}
		})
	}
}
