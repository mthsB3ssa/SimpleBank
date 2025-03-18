package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

// RandomInt generates a random int between min an max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n using a slice of runes.
// This method converts the alphabet into a []rune, making it ready for any Unicode characters.
// It fills a rune slice with random selections from the letters slice, then converts it to a string.
// var letters = []rune("abcdefghijklmnopqrstuvwxyz")
// func RandomString(n int) string {
//     b := make([]rune, n)
//     for i := range b {
//         b[i] = letters[rand.Intn(len(letters))]
//     }
//     return string(b)
// }

// RandomString generates a random string of length n using a strings.Builder.
// This method treats the alphabet as a simple string of bytes (ASCII only) for efficiency.
// It picks a random byte from the alphabet for each position and appends it to the builder.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(1, 1000)
}
