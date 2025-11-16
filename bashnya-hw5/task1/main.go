package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	for _, n := range nums {
		go func(n int) {
			ch <- n * n
		}(n)
	}

	sm := 0
	for range nums {
		sm += <-ch
	}

	fmt.Println("Сумма:", sm)
}
