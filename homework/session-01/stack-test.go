package main

import (
	"fmt"
	"thanhtran/homework/session-01/stack"
)

func main() {
	stack := stack.New()
	fmt.Println(stack)

	// attempt to push
	stack.Push("(")
	stack.Push("{")
	stack.Push("[")
	fmt.Println(stack)
	fmt.Println(stack.Len())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
}
