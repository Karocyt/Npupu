package main

import "fmt"

// SortedHashedTree is the root struct of our tree
type SortedHashedTree struct {
	header       *Node
	length       int64
	dict         map[string]*Node
	history      map[string]bool
	inputsCount  int64
	outputsCount int64
	maxSize      int64
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

func (tree *SortedHashedTree) insert(key string, val interface{}, score float32) bool {
	fmt.Println("Header:", tree.header)
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
		current := next
		if node.score < current.score {
			next = current.left
		} else {
			next = current.right
		}
	}
	if node.score < current.score {
		current.left = node
	} else {
		current.right = node
	}
}

func (tree *SortedHashedTree) delete(key string) bool {
	node := tree.dict[key]
	if node == nil {
		return false
	}
	delete(tree.dict, key)
	if node.parent != nil {
		if node.parent.left == node {
			node.parent.left = node.left
			tree.insertNode(node.right)
		} else if node.parent.right == node {
			node.parent.right = node.left
			tree.insertNode(node.right)
		}
	} else {
		tree.header = node.left
		tree.insertNode(node.right)
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

// GetMin gives you the element with the lowest score value
func (tree *SortedHashedTree) fetchMin() *Node {
	current := tree.header
	for current != nil {
		if current.left != nil {
			current = current.left
		} else {
			return current
		}
	}
	return nil
}

// GetMax gives you the element with the highest score value
func (tree *SortedHashedTree) fetchMax() *Node {
	current := tree.header
	for current != nil {
		if current.right != nil {
			current = current.right
		} else {
			return current
		}
	}
	return nil
}
