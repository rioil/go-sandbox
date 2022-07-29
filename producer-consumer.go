package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	defer close(ch)

	wg.Add(2)

	/*
	 * MEMO
	 * wgを値渡しにするとコピーされたwgのカウンタが更新されることになり，
	 * main関数のWaitは永遠に完了しないためdeadlockが発生する
	 */
	go produce(ch, &wg)
	go consume(ch, &wg)

	wg.Wait()

	fmt.Println("main finished")
}

func produce(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
	}
	ch <- -1

	fmt.Println("produce finished")
}

func consume(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		num := <-ch
		if num == -1 {
			break
		}

		fmt.Println(num * num)
	}

	fmt.Println("consume finished")
}
