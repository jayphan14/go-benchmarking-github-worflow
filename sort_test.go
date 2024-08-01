package main

import (
	"testing"
)

func BenchmarkMySort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := generateArray(10000)
		MySort(arr)
	}
}
