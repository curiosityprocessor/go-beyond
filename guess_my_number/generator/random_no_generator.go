package generator

import (
	"math/rand"
	"time"
)

const MaxValue = 100

func RandomNoGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MaxValue)
}