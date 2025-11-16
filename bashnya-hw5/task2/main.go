package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: завершение по контексту\n", id)
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d: канал закрыт, выходим\n", id)
				return
			}
			fmt.Printf("Worker %d обработал: %d\n", id, v)
		}
	}
}

func main() {
	var n int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&n)

	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	go func() {
		val := 1
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- val
				val++
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	<-sigChan
	fmt.Println("\nПолучен Ctrl+C, завершаем...")

	cancel()
	wg.Wait()

	fmt.Println("Все воркеры завершены. Программа завершена.")
}
