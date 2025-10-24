package main

import (
	"bashnya-hw3/stack"
	"fmt"
)

func main() {

	// Stack
	// ------------------------------- //

	st := stack.New()
	st.Push(999)
	fmt.Println("Size", st.Size())
	fmt.Println("Is empty:", st.IsEmpty())
	fmt.Println("Pop item", st.Pop())
	fmt.Println("Size", st.Size())
	fmt.Println("Is empty:", st.IsEmpty())

	// ------------------------------- //
}
