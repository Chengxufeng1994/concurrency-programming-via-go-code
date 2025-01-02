package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

var (
	maxWorkers int
	sema       *semaphore.Weighted
	task       []int
)

func init() {
	log.Println("init")

	maxWorkers = runtime.GOMAXPROCS(0)
	sema = semaphore.NewWeighted(int64(maxWorkers))
	task = make([]int, maxWorkers*4)
}

func main() {
	ctx := context.Background()

	for i := range task {
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}

		go func(i int) {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond)
			task[i] = i + 1
		}(i)
	}

	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("err: %v", err)
	}

	log.Println(task)
}
