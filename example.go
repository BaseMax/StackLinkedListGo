package main

import "fmt"

func main() {
	// Create a new stack
	stack := NewStackLinkedList()

	// Push some elements
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Print the stack
	stack.Print()

	// Pop the top element
	fmt.Println(stack.Pop())

	// Print the stack
	stack.Print()

	// Reverse the stack
	stack.Reverse()

	// Print the stack
	stack.Print()

	// Reverse the stack (recursive)
	stack.ReverseRecursive(stack.PeekNode())

	// Print the stack
	stack.Print()

	// Get the top element
	fmt.Println(stack.Peek())

	// Get the size of the stack
	fmt.Println(stack.Size())

	// Check if the stack is empty
	fmt.Println(stack.IsEmpty())

	// Clear the stack
	stack.Clear()

	// Check if the stack is empty
	fmt.Println(stack.IsEmpty())

	// Print the stack
	stack.Print()
}
