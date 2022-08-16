package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	const (
		min int64 = 5
		max int64 = 30
	)

	value := RandomInt(min, max)

	require.GreaterOrEqual(t, value, min)
	require.LessOrEqual(t, value, max)
}

func TestRandomString(t *testing.T) {
	length := 8
	randomString := RandomString(length)
	require.Equal(t, len(randomString), length)
}

func TestRandomOwner(t *testing.T) {
	randomOwner := RandomOwner()
	require.Equal(t, len(randomOwner), 6)
}

func TestRandomBalance(t *testing.T) {
	randomBalance := RandomBalance()
	require.GreaterOrEqual(t, randomBalance, int64(0))
}

func TestRandomCurrency(t *testing.T) {
	randomCurrency := RandomCurrency()
	require.Contains(t, []string{"USD", "EUR", "CAD", "AUD"}, randomCurrency)
}
