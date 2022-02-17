package generator

import (
	"math/rand"
	"time"
)

const MaxValue = 100

func GenerateRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MaxValue)
}