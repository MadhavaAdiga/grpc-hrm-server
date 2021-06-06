package utils

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomString(n int) string {
	var builder strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(k)]
		builder.WriteByte(ch)
	}

	return builder.String()
}

func RandomName() string {
	return RandomString(8)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomContactNum() int {
	var builder strings.Builder

	for i := 11; i != 0; i-- {
		ch := fmt.Sprint(RandomInt(0, 9))
		builder.WriteString(ch)
	}
	str := builder.String()

	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("unable to create random contact number: %v", err)
		return -1
	}

	return val
}
