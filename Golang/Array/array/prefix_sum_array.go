package main

import "log"

// 前缀和数组 预处理结构
func main() {
	arr := []int{2, 3, 1, 5, -2}
	log.Println(arr)
	sum := preSumArray(arr)

	l := 1
	r := 3
	log.Println("数组范围[", l, "...", r, "]上累加和:", getSum(sum, l, r))
}

func preSumArray(arr []int) []int {
	n := len(arr)
	sum := make([]int, n)
	sum[0] = arr[0]
	for i := 1; i < n; i++ {
		// 0 ~ i范围上的前缀和, sum[i]
		// 0 ~ i-1范围上的前缀和 + arr[i] -> 0 ~ i范围上的前缀和
		sum[i] = sum[i-1] + arr[i]
	}
	return sum
}

// 原始数组arr[l...r]范围上的累加和返回
func getSum(sum []int, l, r int) int {
	if l == 0 {
		return sum[r]
	} else {
		return sum[r] - sum[l-1]
	}
}
