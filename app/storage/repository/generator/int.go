package generator

import (
	"sync"
	"sync/atomic"
)

type Generate interface {
	Get() int32
	GetSlice(n int32) []int32
	Reset()
	Close()
}

var (
	generators = make(map[interface{}]Generate)
	lock       sync.RWMutex
)

func Get(key interface{}) Generate {
	lock.RLock()
	g, ok := generators[key]
	lock.RUnlock()

	if ok == true {
		return g
	}

	g = newIntGenerator(key)

	lock.Lock()
	generators[key] = g
	lock.Unlock()

	return g
}

type intGenerator struct {
	v   int32
	key interface{}
}

func (g *intGenerator) Get() int32 {
	return atomic.AddInt32(&g.v, step)
}

func (g *intGenerator) GetSlice(n int32) []int32 {
	end := atomic.AddInt32(&g.v, n)
	res := make([]int32, n)
	for i := range res {
		res[i] = end + int32(i)
	}

	return res
}

func (g *intGenerator) Close() {
	lock.Lock()
	delete(generators, g.key)
	lock.Unlock()
}

func (g *intGenerator) Reset() {
	atomic.StoreInt32(&g.v, preStartValue)
}

const (
	preStartValue = -1
	step          = 1
)

func newIntGenerator(key interface{}) *intGenerator {
	return &intGenerator{
		preStartValue,
		key,
	}
}
