package generator

import (
	"crypto/rand"
	"fmt"
)

type CodeGenerator struct{}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

func (cg *CodeGenerator) GenerateCode(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("invalid code length: %d", n)
	}

	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)

	// use crypto/rand for secure randomness
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// map bytes to alphabet
	for i := 0; i < n; i++ {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}
	return string(b), nil
}
