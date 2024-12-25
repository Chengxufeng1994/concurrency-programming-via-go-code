package ch3

import (
	"sync"
	"testing"
	"time"
)

func TestReentrant_Lock(t *testing.T) {
	var mu sync.RWMutex

	mu.Lock()
	{
		mu.Lock()
		t.Log("can not achieve")
		mu.Unlock()
	}
	mu.Unlock()
}

func TestReentrant_RLock(t *testing.T) {
	var mu sync.RWMutex

	mu.RLock()
	{
		mu.RLock()
		t.Log("achieve")
		mu.RUnlock()
	}
	mu.RUnlock()
}

func TestReentrant_DeadLock(t *testing.T) {
	var mu sync.RWMutex

	go func() {
		mu.RLock()
		{
			time.Sleep(10 * time.Second)
			mu.RLock()
			t.Log("can not achieve")
			mu.RUnlock()
		}
		mu.RUnlock()
	}()

	time.Sleep(1 * time.Second)
	mu.Lock()
	t.Log("can not achieve")
	mu.Unlock()
}

func TestReentrant_DeadLock2(t *testing.T) {
	var mu sync.RWMutex

	mu.RLock()
	{
		mu.Lock()
		t.Log("can not achieve")
		mu.Unlock()
	}
	mu.RUnlock()
}

func TestReentrant_DeadLock3(t *testing.T) {
	var mu sync.RWMutex

	mu.Lock()
	{
		mu.RLock()
		t.Log("can not achieve")
		mu.RUnlock()
	}
	mu.Unlock()
}
