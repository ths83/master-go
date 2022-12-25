package dice

import (
	"math/rand"
	"time"
)

func Roll(sides int) int {
	return rand.Intn(sides) + 1
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Seed(n int64) {
	if n == 0 {
		n = time.Now().UnixNano()
	}
	rand.Seed(n)
}
