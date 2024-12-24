package transfer

import (
	"sync"
	"sync/atomic"
)

type Account struct {
	Balance int64
	InTx    bool
}

func transfer1(from, to *Account, amount int64) bool {
	if from.Balance < amount {
		return false
	}

	from.Balance -= amount
	to.Balance += amount
	return true
}

func transfer2(from, to *Account, amount int64) bool {
	bal := atomic.LoadInt64(&from.Balance)
	if bal < amount {
		return false
	}

	atomic.AddInt64(&from.Balance, -amount)
	atomic.AddInt64(&to.Balance, amount)

	return true
}

var mu sync.Mutex

func transfer3(from, to *Account, amount int64) bool {
	mu.Lock()
	defer mu.Unlock()

	bal := atomic.LoadInt64(&from.Balance)
	if bal < amount {
		return false
	}

	atomic.AddInt64(&from.Balance, -amount)
	atomic.AddInt64(&to.Balance, amount)

	return true
}

func transfer4(from, to *Account, amount int64) bool {
	from.InTx = true
	to.InTx = true
	defer func() {
		from.InTx = false
		to.InTx = false
	}()

	mu.Lock()
	defer mu.Unlock()

	bal := atomic.LoadInt64(&from.Balance)
	if bal < amount {
		return false
	}

	atomic.AddInt64(&from.Balance, -amount)
	atomic.AddInt64(&to.Balance, amount)

	return true
}
