package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	assert.Equal(t, 81, square(9), "81이 나와야 하는데")
	assert.Equal(t, 9, square(3), "9가 나와야 하는데")
	assert.Equal(t, 1, square(1), "1이 나와야 하는데")
	assert.Equal(t, 0, square(0), "0이 나와야 하는데")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}

func BenchmarkFibonacci3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci3(20)
	}
}