package main

import (
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	var mu sync.Mutex
	var count int64
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			// defer mu.Unlock()
			count++
		}()
	}

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
