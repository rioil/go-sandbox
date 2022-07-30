package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Print("Input number: ")

		select {
		case input, ok := <-readInput():
			if !ok {
				continue
			}
			fmt.Printf("Waiting %d seconds...\n", input)
			time.Sleep(time.Duration(input) * time.Second)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
			goto end
		}
	}

end:

	fmt.Println("Completed")
}

func readInput() <-chan int {
	ch := make(chan int)

	go func() {
		var second int
		fmt.Scanf("%d\n", &second)
		ch <- second
		close(ch)
	}()

	return ch
}
