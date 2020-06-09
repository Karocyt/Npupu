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
		dict:    make(map[string]*Node, 1000),
		history: make(map[string]bool, 1000),
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
			color:  RED,
			parent: nil,
		}
		tree.dict[key] = &node
		tree.inputsCount++
		return tree.insertNode(&node)
	}
	return false
}

func (tree *SortedHashedTree) insertNode(node *Node) bool {
	if node == nil {
		return false
	}
	tree.length++
	if tree.length > tree.maxSize {
		tree.maxSize = tree.length
	}
	if tree.header == nil {
		tree.header = node
		tree.enforceRB(node)
		return true
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
	tree.enforceRB(node)
	return true
}

func setHeadBlack(tree *SortedHashedTree) {
	if tree.header != nil {
		tree.header.color = BLACK
	}
}

// Delete deletes
func (tree *SortedHashedTree) Delete(key string) bool {
	defer setHeadBlack(tree)
	node := tree.dict[key]
	if node == nil {
		return false
	}
	tree.length--
	delete(tree.dict, key)

	if node.left == nil && node.right == nil {
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.left = nil
			} else {
				node.parent.right = nil
			}
			tree.enforceRB(node.parent)
			return true
		}
		tree.header = nil // else
		return true
	}
	if node.left == nil {
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.left = node.right
			} else {
				node.parent.right = node.right
			}
		} else {
			tree.header = node.right
		}
		node.right.parent = node.parent
		tree.enforceRB(node.parent)
		return true
	}
	if node.right == nil {
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.left = node.left
			} else {
				node.parent.right = node.left
			}
		} else {
			tree.header = node.left
		}
		node.left.parent = node.parent
		tree.enforceRB(node.parent)
		return true
	}
	// if 2 childs
	replacement := getMin(node.right)
	if replacement == nil {
		return false
	}

	if replacement.parent.right == replacement {
		replacement.parent.right = replacement.right
	} else {
		replacement.parent.left = replacement.right
	}
	if replacement.right != nil {
		replacement.right.parent = replacement.parent
	}
	node.key, node.score, node.Value = replacement.key, replacement.score, replacement.Value
	tree.dict[node.key] = node

	//tree.enforceRB(node.parent)
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

func getMin(current *Node) *Node {
	for current != nil {
		if current.left != nil {
			current = current.left
		} else {
			return current
		}
	}
	return nil
}

// GetMin gives you the element with the lowest score value
func (tree *SortedHashedTree) GetMin() interface{} {
	ret := getMin(tree.header).Value
	if ret != nil {
		tree.outputsCount++
	}
	return ret
}

func getMax(current *Node) *Node {
	for current != nil {
		if current.right != nil {
			current = current.right
		} else {
			return current
		}
	}
	return nil
}

// GetMax gives you the element with the highest score value
func (tree *SortedHashedTree) GetMax() interface{} {
	return getMax(tree.header).Value
}
