package main

import (
	"fmt"
	"time"
)

type Token struct{}

const numOfWorker = 4

type Worker struct{}

func newWorker(idx int, ch chan Token, next chan Token) {
	for {
		token := <-ch
		fmt.Println(idx)
		time.Sleep(1 * time.Second)
		next <- token
	}
}

func main() {
	chs := make([]chan Token, numOfWorker)
	for i := 0; i < numOfWorker; i++ {
		chs[i] = make(chan Token)
	}

	for i := 0; i < numOfWorker; i++ {
		go newWorker(i, chs[i], chs[(i+1)%numOfWorker])
	}

	chs[0] <- Token{}

	for {
	}
}
