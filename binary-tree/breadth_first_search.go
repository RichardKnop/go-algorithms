package binarytree

func (t *binaryTree) BreadthFirstSearch(value interface{}) *Node {

	if t.root == nil {
		return nil
	}

	return t.breadthFirstSearch(value, t.root)
}

func (t *binaryTree) breadthFirstSearch(value interface{}, nodes ...*Node) *Node {

	var nodesBellow []*Node
	for _, node := range nodes {

		t.traversed++

		if node.Value == value {
			return node
		}

		nodesBellow = append(nodesBellow, node.Children()...)
	}

	return t.breadthFirstSearch(value, nodesBellow...)
}
