# Stack Linked-List Go

This is a simple implementation of a stack using a linked-list in Go.

## Functions


**Public functions:**

- `func NewStackLinkedList() *StackLinkedList`: Create a new stack

**Member methods:**


**Types:**

```
type Node struct {
	data interface{}
	next *Node
}

type StackLinkedList struct {
	top  *Node
	size int
}
```
