package ch3

import (
	"sync"
	"testing"
)

func BenchmarkCounter_Mutex(b *testing.B) {
	var counter int64
	var mu sync.Mutex

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.RunParallel(func(p *testing.PB) {
			i := 0
			for p.Next() {
				i++

				if i%10000 == 0 {
					mu.Lock()
					counter++
					mu.Unlock()
				} else {
					mu.Lock()
					_ = counter
					mu.Unlock()
				}
			}
		})
	}
}

func BenchmarkCounter_RWMutex(b *testing.B) {
	var counter int64
	var mu sync.RWMutex

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.RunParallel(func(p *testing.PB) {
			i := 0
			for p.Next() {
				i++

				if i%10000 == 0 {
					mu.Lock()
					counter++
					mu.Unlock()
				} else {
					mu.RLock()
					_ = counter
					mu.RUnlock()
				}
			}
		})
	}
}
