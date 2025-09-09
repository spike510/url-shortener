package storage

import (
	"testing"
)

func TestInMemoryStorage_SaveAndGet(t *testing.T) {
	storage := NewInMemoryStorage()

	code := "abc123"
	url := "https://example.com"

	// Test saving a new code-url pair
	err := storage.Save(code, url)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test retrieving the saved URL
	retrievedURL, err := storage.Get(code)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if retrievedURL != url {
		t.Errorf("Expected URL %s, got %s", url, retrievedURL)
	}

	// Test saving a duplicate code
	err = storage.Save(code, "https://another.com")
	if err == nil {
		t.Fatal("Expected error for duplicate code, got nil")
	}

	// Test retrieving a non-existent code
	_, err = storage.Get("nonexistent")
	if err == nil {
		t.Fatal("Expected error for non-existent code, got nil")
	}
}

func TestInMemoryStorage_Empty(t *testing.T) {
	storage := NewInMemoryStorage()

	// Test retrieving from an empty storage
	_, err := storage.Get("anycode")
	if err == nil {
		t.Fatal("Expected error for non-existent code in empty storage, got nil")
	}
}

func TestInMemoryStorage_MultipleEntries(t *testing.T) {
	storage := NewInMemoryStorage()

	entries := map[string]string{
		"code1": "https://example1.com",
		"code2": "https://example2.com",
		"code3": "https://example3.com",
	}

	// Save multiple entries
	for code, url := range entries {
		err := storage.Save(code, url)
		if err != nil {
			t.Fatalf("Expected no error saving %s, got %v", code, err)
		}
	}

	// Retrieve and verify each entry
	for code, expectedURL := range entries {
		retrievedURL, err := storage.Get(code)
		if err != nil {
			t.Fatalf("Expected no error retrieving %s, got %v", code, err)
		}
		if retrievedURL != expectedURL {
			t.Errorf("Expected URL %s for code %s, got %s", expectedURL, code, retrievedURL)
		}

	}
}
