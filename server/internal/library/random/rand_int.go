package random

import (
	"math/rand"
	"time"
)

func GenerateRandomInt64() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63()
}
