package sortedhashedtree

func (tree *SortedHashedTree) rotateLeft(node *Node) {
	nodeRight := node.right
	if nodeRight == nil {
		return
	}
	node.right = nodeRight.left
	if node.right != nil {
		node.right.parent = node
	}
	nodeRight.parent = node.parent
	if node.parent == nil {
		tree.header = nodeRight
	} else if node == node.parent.left {
		node.parent.left = nodeRight
	} else {
		node.parent.right = nodeRight
	}
	nodeRight.left = node
	node.parent = nodeRight
}

func (tree *SortedHashedTree) rotateRight(node *Node) {
	nodeLeft := node.left
	if nodeLeft == nil {
		return
	}
	node.left = nodeLeft.right
	if node.left != nil {
		node.left.parent = node
	}
	nodeLeft.parent = node.parent
	if node.parent == nil {
		tree.header = nodeLeft
	} else if node == node.parent.right {
		node.parent.right = nodeLeft
	} else {
		node.parent.left = nodeLeft
	}
	nodeLeft.right = node
	node.parent = nodeLeft
}

func (tree *SortedHashedTree) enforceRB(node *Node) {
	var p *Node  // parent
	var gp *Node // grand parent

	if node == nil {
		return
	}

	// Remontada while oldies
	for node != tree.header && node.color != BLACK && node.parent.color == RED {

		//fmt.Print("Coucou")
		p = node.parent
		gp = p.parent

		// if p is left child of gp
		if p == gp.left {
			uncle := gp.right

			// if uncle RED
			if uncle != nil && uncle.color == RED {
				gp.color = RED
				p.color = BLACK
				uncle.color = BLACK
				node = gp
			} else {
				if node == p.right { // on le passe à gauche
					tree.rotateLeft(p)
					node = p
					p = p.parent
				}
				// node is left child of its parent
				tree.rotateRight(gp) //// du vieux node ?!
				p.color, gp.color = gp.color, p.color
				node = p

			}
		} else {
			// node is left child, all the same but the other way
			uncle := gp.left

			if uncle != nil && uncle.color == RED {
				gp.color = RED
				p.color = BLACK
				uncle.color = BLACK
				node = gp
			} else {
				if node == p.left {
					tree.rotateRight(p)
					node = p
					p = p.parent
				}
				tree.rotateLeft(gp)
				p.color, gp.color = gp.color, p.color
				node = p

			}
		}
	}
	tree.header.color = BLACK
}
