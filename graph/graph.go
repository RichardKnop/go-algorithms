package graph

import (
	"errors"
	"fmt"
)

type Vertex struct {
	Key      int
	Vertices map[int]*Vertex
}

func NewVertex(key int) *Vertex {

	return &Vertex{
		Key:      key,
		Vertices: make(map[int]*Vertex, 0),
	}
}

type Graph struct {
	Vertices map[int]*Vertex
	directed bool
}

func NewDirectedGraph() *Graph {

	return &Graph{
		Vertices: make(map[int]*Vertex, 0),
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {

	return &Graph{
		Vertices: make(map[int]*Vertex, 0),
	}
}

func (g *Graph) AddVertex(key int) {

	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2 int) error {

	if k1 == k2 {
		return errors.New("can only add edge between unique nodes")
	}

	v1, ok := g.Vertices[k1]
	if !ok {
		return fmt.Errorf("vertex with key %d not found", k1)
	}
	v2, ok := g.Vertices[k2]
	if !ok {
		return fmt.Errorf("vertex with key %d not found", k2)
	}

	// Vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return nil
	}

	// Add a directed edge between v1 and v2
	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	v1.Vertices[v2.Key] = v2
	if !g.directed {
		v2.Vertices[v1.Key] = v1
	}

	return nil
}
