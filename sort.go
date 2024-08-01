package main

import (
	"math/rand"
)

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
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

	return arr
}

func generateArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = 100000 - i
	}
	return arr
}

func MySort(arr []int) []int {
	return QuickSort(arr)
}
