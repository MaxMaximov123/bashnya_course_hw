package main

import "fmt"

func main() {

	chInput := make(chan int)
	chOutput := make(chan int)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	go func() {
		for _, n := range arr {
			chInput <- n
		}
		close(chInput)
	}()

	go func() {
		for v := range chInput {
			chOutput <- v * 2
		}
		close(chOutput)
	}()

	for result := range chOutput {
		fmt.Println(result)
	}

}
