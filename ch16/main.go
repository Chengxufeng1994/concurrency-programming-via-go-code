package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/marusama/cyclicbarrier"
)

func main() {
	cyclicbarrier := cyclicbarrier.New(10)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				log.Printf("goroutine %d 來到第幾 %d 輪\n", i, j)
				err := cyclicbarrier.Await(context.TODO())
				log.Printf("goroutine %d 突破第幾 %d 輪\n", i, j)
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
