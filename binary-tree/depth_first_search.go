package binarytree

// DepthFirstSearch ...
func DepthFirstSearch(node *Tree, value int) *Tree {
	if node.Value == value {
		return node
	}

	if node.Left != nil {
		if found := DepthFirstSearch(node.Left, value); found != nil {
			return found
		}
	}

	if node.Right != nil {
		if found := DepthFirstSearch(node.Right, value); found != nil {
			return found
		}
	}

	return nil
}
