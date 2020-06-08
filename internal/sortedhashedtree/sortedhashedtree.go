package sortedhashedtree

// SortedHashedTree is the root struct of our tree
type SortedHashedTree struct {
	header       *Node
	length       uint64
	dict         map[string]*Node
	history      map[string]bool
	inputsCount  uint64
	outputsCount uint64
	maxSize      uint64
}

func (tree SortedHashedTree) String() string {
	if tree.header != nil {
		return tree.header.String()
	}
	return "Empty"
}

// New returns a new initialized SortedHashedTree
func New() SortedHashedTree {
	elem := SortedHashedTree{
		dict:    map[string]*Node{},
		history: map[string]bool{},
	}

	return elem
}

// Insert inserts
func (tree *SortedHashedTree) Insert(key string, val interface{}, score float32) bool {
	if tree.history[key] == false {
		tree.history[key] = true
		node := Node{
			key:    key,
			Value:  val,
			score:  score,
			color:  BLACK,
			parent: nil,
		}
		tree.dict[key] = &node
		tree.insertNode(&node)
		return true
	}
	return false
}

func (tree *SortedHashedTree) insertNode(node *Node) {
	if node == nil {
		return
	}
	tree.length++
	if tree.length > tree.maxSize {
		tree.maxSize = tree.length
	}
	if tree.header == nil {
		tree.header = node
		return
	}
	next := tree.header
	current := next
	for next != nil {
		current = next
		if node.score < current.score {
			next = current.left
		} else {
			next = current.right
		}
	}
	node.parent = current
	if node.score < current.score {
		current.left = node
	} else {
		current.right = node
	}
}

// Delete deletes
func (tree *SortedHashedTree) Delete(key string) bool {
	node := tree.dict[key]
	if node == nil {
		return false
	}
	tree.length--
	delete(tree.dict, key)
	if node.parent != nil {
		//fmt.Println("Daddy.")
		if node.parent.left == node {
			//fmt.Println("\tI was left")
			node.parent.left = node.left
			tree.insertNode(node.right)
		} else if node.parent.right == node {
			//fmt.Println("\tI was right")
			node.parent.right = node.left
			tree.insertNode(node.right)
		}
		if node.left != nil {
			node.left.parent = node.parent
		}
		if node.right != nil {
			node.right.parent = node.parent
		}
	} else {
		//fmt.Println("No Daddy")
		if node.right != nil {
			node.right.parent = nil
		}
		tree.header = node.right
		//fmt.Println("\t delete head")
		tree.insertNode(node.left)
		//fmt.Println("\t reinserted left elem")
	}
	return true
}

// IsInHistory tells you if this key was already stored
func (tree *SortedHashedTree) IsInHistory(key string) bool {
	return tree.history[key]
}

// GetByKey gives you a Node given a key
func (tree *SortedHashedTree) GetByKey(key string) interface{} {
	return tree.dict[key].Value
}

// GetLen gives you the number of elements
func (tree *SortedHashedTree) GetLen() uint64 {
	return tree.length
}

// GetMin gives you the element with the lowest score value
func (tree *SortedHashedTree) GetMin() interface{} {
	current := tree.header
	for current != nil {
		if current.left != nil {
			current = current.left
		} else {
			return current.Value
		}
	}
	return nil
}

// GetMax gives you the element with the highest score value
func (tree *SortedHashedTree) GetMax() interface{} {
	current := tree.header
	for current != nil {
		if current.right != nil {
			current = current.right
		} else {
			return current.Value
		}
	}
	return nil
}
