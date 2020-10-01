package binarytree

type binaryTree struct {
	root      *Node
	traversed int
}

type Node struct {
	Left  *Node
	Value interface{}
	Right *Node
}

type BinaryTree interface {
	BreadthFirstSearch(value interface{}) *Node
	DepthFirstSearch(value interface{}) *Node
	Traversed() int
}

func New(root *Node) BinaryTree {

	return &binaryTree{root: root}
}

func NewNode(value interface{}, left, right *Node) *Node {

	return &Node{
		Left:  left,
		Value: value,
		Right: right,
	}
}

func (t *binaryTree) Traversed() int {

	return t.traversed
}

func (n *Node) Children() []*Node {

	var children []*Node

	if n.Left != nil {
		children = append(children, n.Left)
	}
	if n.Right != nil {
		children = append(children, n.Right)
	}

	return children
}
