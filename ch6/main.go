package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}

	return o.doSlow(f)
}

func (o *Once) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		defer func() {
			if err == nil {
				atomic.StoreUint32(&o.done, 1)
			}
		}()
		err = f()
	}
	return err
}

func main() {
	var once Once

	f1 := func() error {
		fmt.Println("f1")
		return errors.New("error from f1")
	}

	once.Do(f1)

	f2 := func() error {
		fmt.Println("f2")
		return nil
	}

	once.Do(f2)
}
