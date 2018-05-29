package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstSearch(t *testing.T) {
	t.Parallel()

	tree := &Tree{
		Value: 8,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 5,
				Left: &Tree{
					Value: 9,
				},
				Right: &Tree{
					Value: 10,
				},
			},
			Right: &Tree{
				Value: 6,
			},
		},
		Right: &Tree{
			Value: 4,
			Left: &Tree{
				Value: 7,
				Left: &Tree{
					Value: 11,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Right: &Tree{
				Value: 8,
			},
		},
	}

	found := BreadthFirstSearch(tree, 12)

	if assert.NotNil(t, found) {
		assert.Equal(t, 12, found.Value)
	}
}
