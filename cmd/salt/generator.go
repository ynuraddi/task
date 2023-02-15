package main

import (
	"math/rand"
	"time"
)

const (
	saltSize = 12
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateRandomSalt(saltSize int) []byte {
	rand.Seed(time.Now().UnixNano())
	salt := make([]byte, saltSize)

	for i := 0; i < saltSize; i++ {
		salt[i] = byte(charset[rand.Intn(len(charset))])
	}

	return salt
}
