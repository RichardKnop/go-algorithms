package binarytree

func (t *binaryTree) BFS(visitCallback func(int) bool) {

	if t.root == nil {
		return
	}

	// Keep track of visited nodes
	visited := make(map[int]bool)

	// Mark root node as visited
	visited[t.root.Key] = true

	// Create a queue and enqueue the root node
	queue := make([]*Node, 1)
	queue[0] = t.root

	// Repeat until queue is empty
	for len(queue) > 0 {
		// Get the first node in the queue
		currentNode := queue[0]

		// Dequeue
		queue = queue[1:]

		visited[currentNode.Key] = true

		if visitCallback(currentNode.Key) {
			return
		}

		for _, node := range []*Node{currentNode.Left, currentNode.Right} {
			if node == nil {
				continue
			}
			_, ok := visited[node.Key]
			// If this node has not been visited yet
			if !ok {
				visited[node.Key] = true

				// Enqueue
				queue = append(queue, node)
			}
		}
	}
}
