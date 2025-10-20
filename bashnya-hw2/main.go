package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	new_n := n

	if n >= 12307 {
		fmt.Println("Введенное число больше или равно 12307")
		return
	}

	for n < 12307 {
		if n%9 == 0 && n%13 == 0 {
			fmt.Println(n)
			fmt.Println("service error")
			break
		} else {
			n++
		}

		if n < 0 {
			n *= -1
		} else if n%7 == 0 {
			n *= 39
		} else if n%9 == 0 {
			n = (n * 13) + 1
			continue
		} else {
			n = (n + 2) * 3
		}

	}

	if !(n%13 == 0 && n%9 == 0) {
		fmt.Printf("В ходе выполнения программы число %v было преобразовано в число %v", new_n, n)
	}

}
