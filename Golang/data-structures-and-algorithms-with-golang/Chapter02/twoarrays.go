package main

import (
	"fmt"
)

func main() {
	var TwoDArray [8][8]int
	TwoDArray[3][6] = 18
	TwoDArray[7][4] = 3
	fmt.Println(TwoDArray)
}
