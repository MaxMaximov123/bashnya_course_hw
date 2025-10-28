package main

import (
	"bashnya-hw3/deque"
	"bashnya-hw3/stack"
	"fmt"
)

func main() {

	// Stack
	// ------------------------------- //

	fmt.Println("Stack functions: \n")
	st := stack.New()
	st.Push(999)
	fmt.Println("Size", st.Size())
	fmt.Println("Is empty:", st.IsEmpty())

	pop_res, is_corrct_pop := st.Pop()
	if is_corrct_pop {
		fmt.Println("Pop item", pop_res)
	} else {
		fmt.Println("Stack is empty, pop error", pop_res)
	}
	fmt.Println("Size", st.Size())
	fmt.Println("Is empty:", st.IsEmpty())

	// ------------------------------- //

	// Deque
	// ------------------------------- //

	fmt.Println("\nDeque functions: \n")
	qe := deque.New()
	qe.PushBack(999)
	qe.PushFront(111)
	fmt.Println("Size", qe.Size())
	fmt.Println("Is empty:", qe.IsEmpty())

	qePopBackRes, qePopBackSuccess := qe.PopBack()
	if qePopBackSuccess {
		fmt.Println("PopBack item", qePopBackRes)
	} else {
		fmt.Println("Stack is empty, pop error", qePopBackRes)
	}

	qePopFrontRes, qePopFrontSuccess := qe.PopBack()
	if qePopFrontSuccess {
		fmt.Println("PopFront item", qePopFrontRes)
	} else {
		fmt.Println("Stack is empty, pop error", qePopFrontRes)
	}

	fmt.Println("Size", qe.Size())
	fmt.Println("Is empty:", qe.IsEmpty())

	// ------------------------------- //
}
