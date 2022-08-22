package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate random owner
func RandomOwner() string {
	return RandomString(6)
}

// Generate random balance
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// Generate random currency
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD, AUD}
	return currencies[rand.Intn(len(currencies))]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
