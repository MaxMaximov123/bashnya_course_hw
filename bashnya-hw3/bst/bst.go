package bst

import "fmt"

type Node struct {
	Value     int
	LeftNode  *Node
	RightNode *Node
}

func New() *BST {
	return &BST{}
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(value int) {
	bst.Root = insertInNode(bst.Root, value)
}

func insertInNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.LeftNode = insertInNode(node.LeftNode, value)

	} else if value > node.Value {
		node.RightNode = insertInNode(node.RightNode, value)
	}

	return node
}

func (bst *BST) Remove(value int) {
	bst.Root = removeNode(bst.Root, value)
}

func removeNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.LeftNode = removeNode(node.LeftNode, value)

	} else if value > node.Value {
		node.RightNode = removeNode(node.RightNode, value)
	} else {
		if node.LeftNode == nil {
			return node.RightNode
		} else if node.RightNode == nil {
			return node.LeftNode
		}

		min := findMin(node.RightNode)
		node.Value = min.Value
		node.RightNode = removeNode(node.RightNode, min.Value)
	}

	return node
}

func findMin(node *Node) *Node {
	for node.LeftNode != nil {
		node = node.LeftNode
	}
	return node
}

func PrintTree(node *Node, level int, sight string) {
	if node == nil {
		return
	}
	PrintTree(node.RightNode, level+1, "^")
	for i := 0; i < level; i++ {
		fmt.Printf("  |  ")
	}

	fmt.Printf("%v%v\n", sight, node.Value)
	PrintTree(node.LeftNode, level+1, "v")
}

func (bst *BST) Depth() int {
	return depth(bst.Root)
}

func depth(node *Node) int {
	if node == nil {
		return 0
	}
	leftDepth := depth(node.LeftNode)
	rightDepth := depth(node.RightNode)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func (bst *BST) Find(value int) bool {
	return findNode(bst.Root, value)
}

func findNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return findNode(node.LeftNode, value)
	} else {
		return findNode(node.RightNode, value)
	}
}
