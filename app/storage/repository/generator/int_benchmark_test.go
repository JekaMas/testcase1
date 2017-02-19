package generator

import (
	"testing"
)

var v int32
func BenchmarkGetNew(b *testing.B) {
	var gs = make([]IGenerate, b.N)

	for i := 0; i < b.N; i++ {
		gs[i] = Get(i)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for _, g := range gs {
		v = g.Get()
	}
}

func BenchmarkGet(b *testing.B) {
	var key = int32(1)

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g := Get(key)
		v = g.Get()
	}
	b.StopTimer()
}

func BenchmarkParallelGetNew(b *testing.B) {
	var v int32

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g := Get(n)
		v = g.Get()
	}
	b.StopTimer()

	if v < 0 {
		b.Fatal("Should be non-negative")
	}
}

func BenchmarkParallelGet(b *testing.B) {
	var (
		v   int32
		key = int32(1)
	)

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g := Get(key)
		v = g.Get()
	}
	b.StopTimer()

	if v < 0 {
		b.Fatal("Should be non-negative")
	}
}
