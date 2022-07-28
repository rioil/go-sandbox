package main

import (
	"fmt"
)

func main() {
	Numbers(1, 2, 3, 4, 5, 4, 2, 1, 0)
	maxIdx, maxVal := Maximum(1, 2, 3, 4, 5)
	fmt.Printf("max value is %d @ %d\n", maxVal, maxIdx)
}

func Numbers(numbers ...int) {
	for i, num := range numbers {
		fmt.Printf("numbers[%d]=%d\n", i, num)
	}
}

func Maximum(numbers ...int) (int, int) {
	var maxVal = 0
	var maxIdx = -1

	for i, num := range numbers {
		if num >= maxVal {
			maxVal = num
			maxIdx = i
		}
	}

	return maxIdx, maxVal
}
