package binarytree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RichardKnop/go-algorithms/binary-tree"
)

func TestDFS(t *testing.T) {
	t.Parallel()

	root := &binarytree.Node{
		Key: 1,
		Left: &binarytree.Node{
			Key: 2,
			Left: &binarytree.Node{
				Key: 3,
				Left: &binarytree.Node{
					Key: 4,
				},
				Right: &binarytree.Node{
					Key: 5,
				},
			},
			Right: &binarytree.Node{
				Key: 6,
			},
		},
		Right: &binarytree.Node{
			Key: 7,
			Left: &binarytree.Node{
				Key: 8,
				Left: &binarytree.Node{
					Key: 9,
				},
				Right: &binarytree.Node{
					Key: 10,
				},
			},
			Right: &binarytree.Node{
				Key: 11,
			},
		},
	}

	tree := binarytree.New(root)

	visitedOrder := []int{}

	visitCallback := func(key int) bool {

		visitedOrder = append(visitedOrder, key)

		return false // so we traverse the whole tree
	}

	tree.DFS(visitCallback)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	assert.Equal(t, expected, visitedOrder)
}
