package main

import "log"

func main() {
	arr := []int{10, 3, 0, -9, 2, 100, -22, 54, 23, 7}
	log.Println(arr)
	selectionSort(arr)
	log.Println(arr)

}

func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 长度 >= 2
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// i ~ n-1范围上,找到最小值
		// 最小值, 和i位置的数交换
		// 不是值, 位置
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
