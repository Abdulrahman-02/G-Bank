package util

import (
	"math/rand"
	"time"
)

// init is called before main
func init() {
	rand.Int63n(time.Now().UnixNano())
}

// RandomInt returns a random number between min and max
func RandomInt(min, max int) int64 {
	return int64(rand.Intn(max-min) + min)
}

// RandomString returns a random string of n characters
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// RandomOwner returns a random string of 6 characters
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney returns a random number between 0 and 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency returns a random string of 3 characters
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
