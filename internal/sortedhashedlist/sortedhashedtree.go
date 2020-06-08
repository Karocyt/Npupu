package sortedhashedlist

// SortedHashedList is the root struct of our tree
type SortedHashedList struct {
	header       *Node
	length       int64
	dict         map[string]*Node
	history      map[string]bool
	inputsCount  int64
	outputsCount int64
	maxSize      int64
}

func (tree SortedHashedList) String() string {
	if tree.header != nil {
		return tree.header.String()
	}
	return "Empty"
}

// New returns a new initialized SortedHashedList
func New() SortedHashedList {
	elem := SortedHashedList{
		dict:    map[string]*Node{},
		history: map[string]bool{},
	}

	return elem
}

// Insert adds a new value to the list
func (tree *SortedHashedList) Insert(key string, val interface{}, score float32) bool {
	if tree.history[key] == false {
		tree.history[key] = true
		node := Node{
			key:   key,
			Value: val,
			score: score,
			next:  nil,
		}
		tree.dict[key] = &node
		tree.insertNode(&node)
		return true
	}
	return false
}

func (tree *SortedHashedList) insertNode(node *Node) {
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
	if node.score < tree.header.score {
		node.next = tree.header
		node.next.prev = node
		tree.header = node
		return
	}
	next := tree.header
	current := next
	for next != nil && next.score < node.score {
		current = next
		next = current.next
	}
	node.prev = current
	node.next = current.next
	if node.next != nil {
		node.next.prev = node
	}
	node.prev.next = node
}

// Delete deletes the node at the given key
func (tree *SortedHashedList) Delete(key string) bool {
	node := tree.dict[key]
	if node == nil {
		return false
	}
	tree.length--
	delete(tree.dict, key)

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		tree.header = node.next
		tree.header.prev = nil
	}
	return true
}

// IsInHistory tells you if this key was already stored
func (tree *SortedHashedList) IsInHistory(key string) bool {
	return tree.history[key]
}

// GetByKey gives you a Node given a key
func (tree *SortedHashedList) GetByKey(key string) interface{} {
	return tree.dict[key].Value
}

// GetMin gives you the element with the lowest score value
func (tree *SortedHashedList) GetMin() interface{} {
	return tree.header.Value
}

// GetLen gives you the number of element in our set
func (tree *SortedHashedList) GetLen() interface{} {
	return tree.length
}
