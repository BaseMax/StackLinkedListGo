package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type StackLinkedList struct {
	top  *Node
	size int
}

/**
 * @summary: Create a new stack
 * @return: A new stack
 */
func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{nil, 0}
}

/**
 * @summary: Push a new element to the stack (to the top)
 * @param: data - The data to be pushed
 */
func (stack *StackLinkedList) Push(data interface{}) {
	stack.top = &Node{data, stack.top}
	stack.size++
}

/**
 * @summary: Push a new element to the stack (to the down)
 * @param: data - The data to be pushed
 */
func (stack *StackLinkedList) PushDown(data interface{}) {
	if stack.size == 0 {
		stack.top = &Node{data, nil}
	} else {
		node := stack.top
		for node.next != nil {
			node = node.next
		}
		node.next = &Node{data, nil}
	}
	stack.size++
}

/**
 * @summary: Pop the top element from the stack (from the top)
 * @return: The top element
 */
func (stack *StackLinkedList) Pop() interface{} {
	if stack.size == 0 {
		return nil
	}

	data := stack.top.data
	stack.top = stack.top.next
	stack.size--
	return data
}

/**
 * @summary: Pop the top element from the stack (from the down)
 * @return: The top element
 */
func (stack *StackLinkedList) PopDown() interface{} {
	if stack.size == 0 {
		return nil
	}

	var data interface{}
	if stack.size == 1 {
		data = stack.top.data
		stack.top = nil
	} else {
		node := stack.top
		for node.next.next != nil {
			node = node.next
		}
		data = node.next.data
		node.next = nil
	}
	stack.size--
	return data
}

/**
 * @summary: Get the top element from the stack (from the top)
 * @return: The top element
 */
func (stack *StackLinkedList) Peek() interface{} {
	if stack.size == 0 {
		return nil
	}

	return stack.top.data
}

/**
 * @summary: Get the top element from the stack (from the down)
 * @return: The top element
 */
func (stack *StackLinkedList) PeekDown() interface{} {
	if stack.size == 0 {
		return nil
	}

	node := stack.top
	for node.next != nil {
		node = node.next
	}
	return node.data
}

/**
 * @summary: Get the size of the stack
 * @return: The size of the stack
 */
func (stack *StackLinkedList) Size() int {
	return stack.size
}

/**
 * @summary: Check if the stack is empty
 * @return: True if the stack is empty, false otherwise
 */
func (stack *StackLinkedList) IsEmpty() bool {
	return stack.size == 0
}

/**
 * @summary: Clear the stack
 */
func (stack *StackLinkedList) Clear() {
	stack.top = nil
	stack.size = 0
}

/**
 * @summary: Print the stack
 */
func (stack *StackLinkedList) Print() {
	for node := stack.top; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}
