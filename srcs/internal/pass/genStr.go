package pass

import (
	"math/rand"
	"strings"
)

func GenStrPass(n int) string {
	var characters = []rune("abcdef0123456789")
	var sb strings.Builder

	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(characters))
		randomChar := characters[randomIndex]
		sb.WriteRune(randomChar)
	}

	return sb.String()
}
