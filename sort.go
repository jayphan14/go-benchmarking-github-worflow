package main

import (
	"math/rand"
	"time"
)

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	pivotIndex := rand.Int() % len(arr)
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}

func generateArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = 100000 - i
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := generateArray(10000)
	BubbleSort(arr)
}
