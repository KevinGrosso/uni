package linkedList

type Node struct {
	Content string
	Next *Node
}

// A List is a pointer to a Node.
type List *Node
