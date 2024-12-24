package ch2

import (
	"sync"
	"testing"
)

func TestOnlyUnLock(t *testing.T) {
	var mu sync.Mutex
	mu.Unlock()
}
