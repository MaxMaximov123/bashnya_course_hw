package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()

			mu.Lock()
			m[x] = x * x
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Результат map:", m)
}
