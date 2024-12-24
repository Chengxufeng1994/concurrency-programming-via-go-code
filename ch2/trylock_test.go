package ch2

import (
	"sync"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	var mu sync.Mutex

	go func() {
		mu.Lock()
		time.Sleep(2 * time.Second)
		mu.Unlock()
	}()

	time.Sleep(1 * time.Second)

	if mu.TryLock() {
		println("tryLock success")
		mu.Unlock()
	} else {
		println("tryLock failed")
	}
}
