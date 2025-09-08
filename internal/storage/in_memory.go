package storage

import "fmt"

type InMemoryStorage struct {
	data map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{data: make(map[string]string)}
}

func (s *InMemoryStorage) Save(code, url string) error {
	if _, exists := s.data[code]; exists {
		return fmt.Errorf("code already exists")
	}

	s.data[code] = url
	return nil
}

func (s *InMemoryStorage) Get(code string) (string, error) {
	url, exists := s.data[code]
	if !exists {
		return "", fmt.Errorf("code not found")
	}
	return url, nil
}
