package generator

import (
	"math/rand"
	"sync"
	"time"
)

type randomGenerator struct {
	r *rand.Rand
	sync.Mutex
}

func getRandomGenerator() *randomGenerator {
	return &randomGenerator{
		rand.New(rand.NewSource(time.Now().UnixNano())),
		sync.Mutex{},
	}
}

func (gen *randomGenerator) Int31n(n int32) int32 {
	gen.Lock()
	res := gen.r.Int31n(n)
	gen.Unlock()

	return res
}

func (gen *randomGenerator) Intn(n int) int {
	gen.Lock()
	res := gen.r.Intn(n)
	gen.Unlock()

	return res
}

func (gen *randomGenerator) setLesserRandom(n *int) {
	*n = gen.Intn(*n) + 1
}
