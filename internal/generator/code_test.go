package generator

import (
	"regexp"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	cg := NewCodeGenerator()

	code, err := cg.GenerateCode(10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(code) != 10 {
		t.Errorf("Expected code length 10, got %d", len(code))
	}

	matched, err := regexp.MatchString("^[0-9a-zA-Z]{10}$", code)
	if err != nil {
		t.Fatalf("Error matching regex: %v", err)
	}
	if !matched {
		t.Errorf("Generated code %s does not match expected pattern", code)
	}
}

func TestGenerateCode_InvalidLength(t *testing.T) {
	cg := NewCodeGenerator()

	_, err := cg.GenerateCode(0)
	if err == nil {
		t.Fatal("Expected error for code length 0, got nil")
	}

	_, err = cg.GenerateCode(-5)
	if err == nil {
		t.Fatal("Expected error for negative code length, got nil")
	}
}
