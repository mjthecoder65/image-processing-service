package utils

import "math/rand"

const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {

	b := make([]byte, length)

	for i := range b {
		b[i] = CHARSET[rand.Intn(len(CHARSET))]
	}

	return string(b)
}
