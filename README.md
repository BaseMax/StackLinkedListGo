# Stack Linked-List Go

This is a simple implementation of a stack using a linked-list in Go.

## Functions

**Public functions:**

- `func NewStackLinkedList() *StackLinkedList`: Create a new stack

**Member methods:**

- `func (stack *StackLinkedList) Push(data interface{})`: Push a new element to the stack (to the top)
- `func (stack *StackLinkedList) PushDown(data interface{})`: Push a new element to the stack (to the down)
- `func (stack *StackLinkedList) Pop() interface{}`: Pop the top element from the stack (from the top)
- `func (stack *StackLinkedList) PopDown() interface{}`: Pop the top element from the stack (from the down)
- `func (stack *StackLinkedList) Peek() interface{}`: Get the top element from the stack (from the top)
- `func (stack *StackLinkedList) PeekDown() interface{}`: Get the top element from the stack (from the down)
- `func (stack *StackLinkedList) Size() int`: Get the size of the stack
- `func (stack *StackLinkedList) IsEmpty() bool`: Check if the stack is empty
- `func (stack *StackLinkedList) Clear()`: Clear the stack
- `func (stack *StackLinkedList) Print()`: Print the stack
- `func (stack *StackLinkedList) Reverse()`: Reverse the stack
- `func (stack *StackLinkedList) ReverseRecursive(Node *Node) *Node`: Reverse the stack recursively

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
