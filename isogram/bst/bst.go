// binary search tree package
package bst

type Tree struct {
	root *Node
}

type Node struct {
	key   byte
	left  *Node
	right *Node
}

// Tree
func (t *Tree) Insert(data byte) int {
	if t.root == nil {
		t.root = &Node{key: data}
		return 1
	} else {
		return t.root.Insert(data)
	}
}

// Node
func (n *Node) Insert(data byte) int {
	if data == n.key {
		return 0
	} else {
		if data < n.key {
			if n.left == nil {
				n.left = &Node{key: data}
				return -1
			} else {
				return n.left.Insert(data)
			}
		} else {
			if n.right == nil {
				n.right = &Node{key: data}
				return 1
			} else {
				return n.right.Insert(data)
			}
		}
	}
}
