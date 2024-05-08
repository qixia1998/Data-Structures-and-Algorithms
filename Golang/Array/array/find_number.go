package main

import "log"

func main() {
	// arr有序
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	search := 10
	log.Println("数组为: ", arr)
	log.Println("搜索的数为: ", search)
	log.Println("是否找到: ", exist(arr, search))
}

// arr数组有序
func exist(arr []int, ele int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	l := 0            // 左边界
	r := len(arr) - 1 // 有边界
	m := 0            // 中点变量

	for l < r {
		// m = (l + r) / 2
		m = l + (r-l)>>1 // 右移相当于/2
		if arr[m] == ele {
			return true
		} else if arr[m] > ele {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
