package main

// 找出数组中唯一的重复元素
import (
	"fmt"
)

// Hash法
// 方法功能：在数组中找唯一重复的元素
// 返回值：重复元素的值，如果无重复元素则返回-1
func FindDupByHash(arr []int) int {
	if arr == nil {
		return -1
	}
	data := map[int]bool{}
	for _, v := range arr {
		if _, ok := data[v]; ok {
			return v
		} else {
			data[v] = true
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 4, 2, 5, 3}
	fmt.Println("Hash 法")
	fmt.Println(FindDupByHash(arr))
}
