package stack

// Node is the item of stack
type Node struct {
	next  *Node
	Value interface{}
}

// NewNode returns Node instance
func NewNode(value interface{}) *Node {
	return &Node{
		Value: value,
	}
}
