package binarytree

func (t *binaryTree) DepthFirstSearch(value interface{}) *Node {

	if t.root == nil {
		return nil
	}

	return t.depthFirstSearch(t.root, value)
}

func (t *binaryTree) depthFirstSearch(node *Node, value interface{}) *Node {

	t.traversed++

	if node.Value == value {
		return node
	}

	if node.Left != nil {
		if found := t.depthFirstSearch(node.Left, value); found != nil {
			return found
		}
	}

	if node.Right != nil {
		if found := t.depthFirstSearch(node.Right, value); found != nil {
			return found
		}
	}

	return nil
}
