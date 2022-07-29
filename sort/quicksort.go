package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 0, -1, 5, 3}
	print(data)

	sort(data, 0, len(data)-1)
	print(data)
}

func sort(data []int, left, right int) {
	if left >= right {
		return
	}

	pivot := data[(left + right) / 2]
	i := left
	j := right
	for {
		for ; i < right && data[i] < pivot; i++ {
		}
		for ; j > left && data[j] > pivot; j-- {
		}

		if i >= j {
			break
		}

		swap(&data[i], &data[j])
		i++
		j--
	}

	print(data)
	sort(data, left, i-1)
	sort(data, j+1, right)
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func print(data []int) {
	for _, val := range data {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
