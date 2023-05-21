package util

import (
	"math/rand"
	"strings"
	"time"
)

var alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	newSeed := time.Now().UnixNano()
	rand.New(rand.NewSource(newSeed))
}

// RandomInt generate a random number
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random String generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner string
func RandomOwner() string {
	return RandomString(6)
}

// Random money amount
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Random currency generator
func RandomCurrency() string {
	currencies := []string{"USD", "GBP", "EUR", "NGN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
