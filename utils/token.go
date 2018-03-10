package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomToken(bytes int) string {
	b := make([]byte, bytes)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
