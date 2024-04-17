package main

import "log"

func main() {
	arr := []int{10, 3, 0, -9, 2, 100, -22, 54, 23, 7}
	log.Println(arr)
	insertionSort(arr)
	log.Println(arr)
}

func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j-1)
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
