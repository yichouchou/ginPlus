package utils

import (
	"testing"
)

func TestContainsHttpMethod(t *testing.T) {
	tests := []struct {
		name    string
		annoDoc string
		expect  string
		isFound bool
	}{
		{"GET", "@GET /test", "GET", true},
		{"POST", "@POST /test", "POST", true},
		{"No method", "some description", "ANY", false},
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
	router, found := ContainsHttpRouter("@GET /hello")
	if router != "/hello" {
		t.Errorf("Expected '/hello', got '%s'", router)
	}
	if !found {
		t.Error("Expected found=true")
	}
}

func TestReverse(t *testing.T) {
	arr := []string{"a", "b", "c"}
	Reverse(&arr)
	expected := []string{"c", "b", "a"}
	for i, v := range expected {
		if arr[i] != v {
			t.Errorf("Expected '%s', got '%s'", v, arr[i])
		}
	}
}

func TestMapMerge(t *testing.T) {
	first := map[string]string{"a": "1"}
	second := map[string]string{"b": "2"}
	result := MapMerge(first, second)
	if len(result) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(result))
	}
}

func TestRandString(t *testing.T) {
	s := RandString(10)
	if len(s) == 0 {
		t.Error("Expected non-empty string")
	}
}
