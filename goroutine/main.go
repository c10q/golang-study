package main

import (
	"fmt"
	"time"
)

func 함수() {
	for i := 0; i < 10; i++ {
		fmt.Println("함수():", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func 또다른함수() {
	for i := 0; i < 10; i++ {
		fmt.Println("또다른함수():", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go 함수()
	go 또다른함수()

	time.Sleep(3 * time.Second)
}
