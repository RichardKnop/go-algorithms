package binarytree

type binaryTree struct {
	root *Node
}

type Node struct {
	Left  *Node
	Key   int
	Right *Node
}

type BinaryTree interface {
	BFS(func(int) bool)
	DFS(func(int) bool)
}

func New(root *Node) BinaryTree {

	return &binaryTree{root: root}
}

func NewNode(key int, left, right *Node) *Node {

	return &Node{
		Left:  left,
		Key:   key,
		Right: right,
	}
}
