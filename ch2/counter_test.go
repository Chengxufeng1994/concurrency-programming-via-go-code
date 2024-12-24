package ch2

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestCounter(t *testing.T) {
	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			for j := 0; j < 100000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	if counter != 64*100000 {
		t.Fatalf("counter = %d, want %d", counter, 64*100000)
	}
}

func TestCounterWithMutex(t *testing.T) {
	var counter int64

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	if counter != 64*100000 {
		t.Fatalf("counter = %d, want %d", counter, 64*100000)
	}
}

func TestCounterWithAtomic(t *testing.T) {
	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	if counter != 64*100000 {
		t.Fatalf("counter = %d, want %d", counter, 64*100000)
	}
}
