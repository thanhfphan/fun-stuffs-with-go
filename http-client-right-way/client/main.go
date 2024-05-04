package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	// n = 100
	n = 1_000
	// n = 10_000 // why it is so random with this ???
)

func main() {
	normal()
	custom()
}

func custom() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConnsPerHost = 10
	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	var wg sync.WaitGroup
	start := time.Now()
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := httpClient.Get("http://127.0.0.1:8080/ping")
			if err != nil {
				fmt.Println("request got err:", err)
				return
			}
			defer resp.Body.Close()
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Time processed in CUSTOM way took", duration.Milliseconds(), "miliseconds")
}

func normal() {
	var defaultClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	var wg sync.WaitGroup
	start := time.Now()
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := defaultClient.Get("http://127.0.0.1:8080/ping")
			if err != nil {
				fmt.Println("request got err:", err)
				return
			}
			defer resp.Body.Close()
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Time processed in NORMAL way took", duration.Milliseconds(), "miliseconds")
}
