package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.baidu.com",
	}
	results := make([]bool, len(urls))
	client := http.Client{
		Timeout: time.Second,
	}

	wg.Add(len(urls))
	for i := 0; i < len(urls); i++ {
		url := urls[i]
		go func(url string) {
			defer wg.Done()

			log.Println("fetching", url)
			resp, err := client.Get(url)
			if err != nil {
				results[i] = false
				return
			}

			results[i] = resp.StatusCode == http.StatusOK
			resp.Body.Close()
		}(url)
	}

	wg.Wait()
	log.Println("done")
	for i := 0; i < len(urls); i++ {
		log.Println(urls[i], ":", results[i])
	}
}
