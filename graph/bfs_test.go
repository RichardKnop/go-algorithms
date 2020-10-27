package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RichardKnop/go-algorithms/graph"
)

func TestBFS(t *testing.T) {
	t.Parallel()

	g := graph.NewDirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 1)

	visitedOrder := []int{}

	visitCallback := func(key int) bool {

		visitedOrder = append(visitedOrder, key)

		return false // so we traverse the whole tree
	}

	graph.BFS(g, 1, visitCallback)

	expected := []int{1, 2, 3, 4}

	assert.Equal(t, expected, visitedOrder)
}
