package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func 함수() {
	for i := 0; i < 10; i++ {
		fmt.Println("함수():", i)
		time.Sleep(500 * time.Millisecond)
	}

	wg.Done()
}

func 또다른함수() {
	for i := 0; i < 10; i++ {
		fmt.Println("또다른함수():", i)
		time.Sleep(400 * time.Millisecond)
	}

	wg.Done()
}

func main() {
	wg.Add(1)
	go 함수()
	wg.Add(1)
	go 또다른함수()

	wg.Wait()

	// time.Sleep(3 * time.Second)
}
