package graph

import (
	"fmt"
)

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

func main() {

	g := NewDirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 1)

	visitCallback := func(key int) bool {

		fmt.Println("Visiting vertex with key ", key)

		return false // so we traverse the whole graph
	}

	BFS(g, 1, visitCallback)
}
