package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomString(6)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomType generates a random type
func RandomType() string {
	types := []string{"topic", "opinion"}
	n := len(types)
	return types[rand.Intn(n)]
}

// RandomCategory generates a random topic category
func RandomCategory() string {
	types := []string{"news", "sport", "entertainment", "covid", "economy", "tech", "fashion", "life", "gourmet", "browse", "culture", "anime", "funny", "love"}
	n := len(types)
	return types[rand.Intn(n)]
}
