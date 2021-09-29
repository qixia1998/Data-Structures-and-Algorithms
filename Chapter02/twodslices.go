package main

import (
	"fmt"
)

func main() {
	var rows int
	var cols int
	rows = 7
	cols = 9
	var twodslices = make([][]int, rows)
	var i int
	for i = range twodslices {
		twodslices[i] = make([]int, cols)
	}
	fmt.Println(twodslices)
}
