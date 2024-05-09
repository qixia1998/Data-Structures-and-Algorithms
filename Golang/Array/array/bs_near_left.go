package main

import "log"

func main() {
	//arr 有序
	arr := []int{1, 5, 9, 13, 25, 46}
	search := 6
	log.Println("数组为: ", arr)
	log.Println(">=", search, "的最左位置为下标: ", moreEqualMostLeft(arr, search))
}

func moreEqualMostLeft(arr []int, num int) int {
	ans := -1
	if arr == nil || len(arr) < 1 {
		return ans
	}
	l := 0
	r := len(arr) - 1
	m := 0
	for l < r {
		m = l + ((r - l) >> 1)
		if arr[m] > num {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}
