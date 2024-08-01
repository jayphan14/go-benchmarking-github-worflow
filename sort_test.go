package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMySort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		arr := generateArray(10000)
		BubbleSort(arr)
	}
}
