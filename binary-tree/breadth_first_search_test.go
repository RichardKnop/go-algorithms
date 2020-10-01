package binarytree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RichardKnop/go-algorithms/binary-tree"
)

func TestBreadthFirstSearch(t *testing.T) {
	t.Parallel()

	root := &binarytree.Node{
		Value: 8,
		Left: &binarytree.Node{
			Value: 2,
			Left: &binarytree.Node{
				Value: 5,
				Left: &binarytree.Node{
					Value: 9,
				},
				Right: &binarytree.Node{
					Value: 10,
				},
			},
			Right: &binarytree.Node{
				Value: 6,
			},
		},
		Right: &binarytree.Node{
			Value: 4,
			Left: &binarytree.Node{
				Value: 7,
				Left: &binarytree.Node{
					Value: 11,
				},
				Right: &binarytree.Node{
					Value: 12,
				},
			},
			Right: &binarytree.Node{
				Value: 8,
			},
		},
	}

	tree := binarytree.New(root)

	found := tree.BreadthFirstSearch(8)

	if assert.NotNil(t, found) {
		assert.Equal(t, 8, found.Value)
		assert.Equal(t, 1, tree.Traversed())
	}

	tree = binarytree.New(root)

	found = tree.BreadthFirstSearch(5)

	if assert.NotNil(t, found) {
		assert.Equal(t, 5, found.Value)
		assert.Equal(t, 4, tree.Traversed())
	}
}
