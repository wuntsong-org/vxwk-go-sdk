package utils

import (
	"math/rand"
	"time"
)

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateUniqueNumber(length int) string {
	const numbers = "0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = numbers[rander.Intn(len(numbers))]
	}

	return string(result)
}

func GenerateRandomText(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789?"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = charset[rander.Intn(len(charset))]
	}

	return string(result)
}
