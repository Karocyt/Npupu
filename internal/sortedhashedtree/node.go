package main

// Colors used to balance the tree
const (
	RED   bool = false // red as default
	BLACK bool = true
)

// Node is the Node struct
type Node struct {
	key    string      // unique key of this node
	Value  interface{} // associated data
	score  float32     // score to determine the order of this node in the set
	color  bool        // Color of the node
	parent *Node       // parent Node
	left   *Node       // Node with lower score
	right  *Node       // Node with higher or equal score
}

// Key Get the key of the node
func (node *Node) Key() string {
	return node.key
}

// Score Get the score of the node
func (node *Node) Score() float32 {
	return node.score
}

// String conforms to the Stringer interface
func (node *Node) String() string {
	ret := ""
	if node.left != nil {
		ret += node.left.String()
	}
	ret += node.key
	if node.right != nil {
		ret += node.right.String()
	}
	return ret
}
