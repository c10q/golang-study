package main

import "fmt"

func square(x int) int {
	return x * x
}

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n - 1) + fibonacci1(n - 2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	fir := 0
	sec := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = fir + sec
		fir = sec
		sec = result
	}
	return result
}

var m map[int]int

func fibonacci3(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	if m == nil {
		m = make(map[int]int)
	}
	if v, ok := m[n]; ok {
		return v
	}

	v := fibonacci3(n - 1) + fibonacci3(n - 2)
	m[n] = v
	return v
}

func main() {
	fmt.Printf("9 x 9 = %d\n", square(9))
	fmt.Printf("rec fun fibo(20): %d\n", fibonacci1(20))
	fmt.Printf("loop fun fibo(20): %d\n", fibonacci2(20))
}