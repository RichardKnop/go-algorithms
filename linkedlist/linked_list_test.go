package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RichardKnop/go-algorithms/linkedlist"
)

func Test_Append(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()

	l.Append(linkedlist.Item("a"))
	l.Append(linkedlist.Item("b"))
	l.Append(linkedlist.Item("c"))

	expected := []string{"a", "b", "c"}

	node := l.Head()
	i := 0
	for node != nil {

		if assert.NotNil(t, node.Value()) {
			assert.Equal(t, expected[i], node.Value())
		}

		node = node.Next()
		i++
	}
}
