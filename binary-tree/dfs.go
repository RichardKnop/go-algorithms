package binarytree

func (t *binaryTree) DFS(visitCallback func(int) bool) {

	if t.root == nil {
		return
	}

	// Keep track of visited nodes
	visited := make(map[int]bool)

	t.dfs(t.root, visited, visitCallback)
}

func (t *binaryTree) dfs(node *Node, visited map[int]bool, visitCallback func(int) bool) {

	// Mark node as visited
	visited[node.Key] = true

	if visitCallback(node.Key) {
		return
	}

	if node.Left != nil {
		_, ok := visited[node.Left.Key]
		if !ok {
			t.dfs(node.Left, visited, visitCallback)
		}
	}

	if node.Right != nil {
		_, ok := visited[node.Right.Key]
		if !ok {
			t.dfs(node.Right, visited, visitCallback)
		}
	}
}
