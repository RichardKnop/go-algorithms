package graph

func BFS(graph *Graph, key int, visitCallback func(int) bool) {

	// Keep track of visited vertices
	visited := make(map[int]bool)

	// Mark starting vertex as visited
	visited[key] = true

	// Create a queue and enqueue the starting vertex
	queue := make([]*Vertex, 1)
	queue[0] = graph.Vertices[key]

	// Repeat until queue is empty
	for len(queue) > 0 {
		// Get the first vertex in the queue
		currentVertex := queue[0]

		// Dequeue
		queue = queue[1:]

		visited[currentVertex.Key] = true

		if visitCallback(currentVertex.Key) {
			return
		}

		for _, vertex := range currentVertex.Vertices {
			_, ok := visited[vertex.Key]
			// If this vertex has not been visited yet
			if !ok {
				visited[vertex.Key] = true

				// Enqueue
				queue = append(queue, vertex)
			}
		}
	}
}
