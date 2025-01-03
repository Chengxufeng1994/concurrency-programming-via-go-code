package ch3

import (
	"net/http"
	"sync"
	"time"
)

type Config struct {
	Group          string
	Retries        int
	ConnectTimeout time.Duration
	IdleTimeout    time.Duration
}

var configMutex sync.RWMutex
var config = &Config{}

func updateConfig(newConfig *Config) {
	configMutex.Lock()
	defer configMutex.Unlock()

	config = newConfig
}

func tryUpdateConfig(newConfig *Config) {
	if ok := configMutex.TryLock(); !ok {
		return
	}

	defer configMutex.Unlock()

	config = newConfig
}

func accessExampleSite() {
	configMutex.RLock()
	retries := config.Retries
	configMutex.RUnlock()

	for i := 0; i < retries; i++ {
		resp, err := http.Get("http://www.example.com")
		if err != nil {
			continue
		}

		resp.Body.Close()
	}
}
