package generator

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func setLesserRandom(n *int) {
	*n = r.Intn(*n) + 1
}
