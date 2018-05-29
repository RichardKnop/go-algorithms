package binarytree

// BreadthFirstSearch ...
func BreadthFirstSearch(node *Tree, value int) *Tree {
	return breadthFirstSearch(value, node)
}

func breadthFirstSearch(value int, nodes ...*Tree) *Tree {
	var children []*Tree
	for _, node := range nodes {
		if node.Value == value {
			return node
		}
		if node.Left != nil {
			children = append(children, node.Left)
		}
		if node.Right != nil {
			children = append(children, node.Right)
		}
	}
	return breadthFirstSearch(value, children...)
}
