package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := readNumbers("numbers.txt")
	if numbers == nil {
		fmt.Println("Error")
		return
	}

	for _, num := range *numbers {
		fmt.Println(num)
	}
}

func readNumbers(path string) *[]int {
	fp, err := os.Open(path)

	if err != nil {
		return nil
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	if !scanner.Scan() {
		return nil
	}
	numCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil
	}

	numbers := make([]int, 0, numCount)

	for i := 0; i < numCount; i++ {
		if scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil
			} else {
				numbers = append(numbers, num)
			}
		}
	}

	return &numbers
}
