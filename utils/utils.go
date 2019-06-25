package utils

import (
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generates a random string of length n
func RandomString(n int) string {
	b := make([]byte, rand.Intn(n)+n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// MaybePanic panics if error is present
func MaybePanic(desc string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %s", desc, err))
	}
}
