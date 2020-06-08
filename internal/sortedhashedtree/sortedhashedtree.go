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
	return true
}

// Delete deletes
func (tree *SortedHashedTree) Delete(key string) bool {
	node := tree.dict[key]
	//fmt.Println("Coucou ?")
	if node == nil {
		// panic(errors.New("Cata"))
		return false
	}
	//fmt.Println("Coucou !")
	tree.length--
	delete(tree.dict, key)

	if node.left == nil && node.right == nil {
		//fmt.Println("Coucou orphan")
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.left = nil
				return true
			} else if node.parent.right == node {
				node.parent.right = nil
				return true
			}
			//fmt.Println("ERROR: Please be kind on us.")
			return false
		}
		tree.header = nil
		return true
	}
	if node.left == nil {
		//fmt.Println("Coucou l guy")
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
		return true
	}
	if node.right == nil {
		//fmt.Println("Coucou r guy")
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
	}

	//fmt.Println("Deleting", node.key)
	replacement := getMin(node.right)
	if replacement == nil {
		return false
	}
	if replacement.parent != node {
		replacement.parent.left = nil
	} else {
		replacement.parent.right = nil
	}
	replacement.left = node.left

	node.key, node.score, node.Value, node.right = replacement.key, replacement.score, replacement.Value, replacement.right
	tree.dict[node.key] = node

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
	return getMin(tree.header).Value
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
