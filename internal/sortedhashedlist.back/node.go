package sortedhashedlist

// Node is the Node struct
type Node struct {
	key   string      // unique key of this node
	Value interface{} // associated data
	score float32     // score to determine the order of this node in the set
	next  *Node       // Node with higher or equal score
	prev  *Node       // Node with lower or equal score
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
	ret := node.key
	if node.next != nil {
		ret += node.next.String()
	}
	return ret
}
