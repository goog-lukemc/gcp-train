package main

// The prize goes to the one who fixes this code.
// All goroutines must exit before main exits.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var hit *time.Ticker

func main() {
	ctx, quit := context.WithTimeout(context.Background(), time.Second*10)
	hit = time.NewTicker(time.Millisecond * 500)

	b := make(chan string)
	wg.Add(2)
	go func() {
		go ping(ctx, quit, b)
		go pong(ctx, quit, b)
	}()

	fmt.Println("Back on Main thread")

	// Lets serve
	b <- "ping"
	<-ctx.Done()
	hit.Stop()
	wg.Wait()
}

func ping(ctx context.Context, cancel context.CancelFunc, bus chan string) {
	for range hit.C {
		select {
		case v := <-bus:
			if v == "ping" {
				fmt.Printf("%s\n", v)
				time.Sleep(time.Millisecond * 500)
				bus <- "pong"
			} else {
				bus <- v
			}
		}

	}
	wg.Done()
	fmt.Println("Exiting Ping")
}

func pong(ctx context.Context, cancel context.CancelFunc, bus chan string) {
	for range hit.C {
		select {
		case v := <-bus:
			if v == "pong" {
				fmt.Printf("%s\n", v)
				time.Sleep(time.Millisecond * 500)
				bus <- "ping"
			} else {
				bus <- v
			}
		}
	}
	wg.Done()
	fmt.Println("Exiting Pong")
}
