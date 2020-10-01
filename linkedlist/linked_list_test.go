package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RichardKnop/go-algorithms/linkedlist"
)

func Test_Empty(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()

	assert.True(t, l.IsEmpty())

	l.Append(linkedlist.Item("a"))

	assert.False(t, l.IsEmpty())
}

func Test_Append(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()

	l.Append(linkedlist.Item("a"))
	l.Append(linkedlist.Item("b"))
	l.Append(linkedlist.Item("c"))

	expected := []string{"a", "b", "c"}
	assertLinkedList(t, expected, l)
}

func Test_Insert_EmptyList_Error(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()

	err := l.Insert(10, linkedlist.Item("a"))

	assert.Error(t, err)
}

func Test_Insert_EmptyList(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()

	err := l.Insert(0, linkedlist.Item("a"))

	if assert.NoError(t, err) {
		expected := []string{"a"}
		assertLinkedList(t, expected, l)
	}
}

func Test_Insert_Beginning(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()
	l.Append(linkedlist.Item("a"))
	l.Append(linkedlist.Item("b"))
	l.Append(linkedlist.Item("c"))

	err := l.Insert(0, linkedlist.Item("hello"))

	if assert.NoError(t, err) {
		expected := []string{"hello", "a", "b", "c"}
		assertLinkedList(t, expected, l)
	}
}

func Test_Insert_Middle(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()
	l.Append(linkedlist.Item("a"))
	l.Append(linkedlist.Item("b"))
	l.Append(linkedlist.Item("c"))

	err := l.Insert(2, linkedlist.Item("hello"))

	if assert.NoError(t, err) {
		expected := []string{"a", "b", "hello", "c"}
		assertLinkedList(t, expected, l)
	}
}

func Test_Insert_End(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()
	l.Append(linkedlist.Item("a"))
	l.Append(linkedlist.Item("b"))
	l.Append(linkedlist.Item("c"))

	err := l.Insert(3, linkedlist.Item("hello"))

	if assert.NoError(t, err) {
		expected := []string{"a", "b", "c", "hello"}
		assertLinkedList(t, expected, l)
	}
}

func Test_RemoveAt_EmptyList_Error(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()

	item, err := l.RemoveAt(0)

	assert.Nil(t, item)
	assert.Error(t, err)
}

func Test_RemoveAt_Beginning(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()
	l.Append(linkedlist.Item("a"))
	l.Append(linkedlist.Item("b"))
	l.Append(linkedlist.Item("c"))

	item, err := l.RemoveAt(0)

	if assert.NoError(t, err) && assert.NotNil(t, item) {
		assert.Equal(t, "a", item)

		expected := []string{"b", "c"}
		assertLinkedList(t, expected, l)
	}
}

func Test_RemoveAt_Middle(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()
	l.Append(linkedlist.Item("a"))
	l.Append(linkedlist.Item("b"))
	l.Append(linkedlist.Item("c"))

	item, err := l.RemoveAt(1)

	if assert.NoError(t, err) && assert.NotNil(t, item) {
		assert.Equal(t, "b", item)

		expected := []string{"a", "c"}
		assertLinkedList(t, expected, l)
	}
}

func Test_RemoveAt_End(t *testing.T) {
	t.Parallel()

	l := linkedlist.New()
	l.Append(linkedlist.Item("a"))
	l.Append(linkedlist.Item("b"))
	l.Append(linkedlist.Item("c"))

	item, err := l.RemoveAt(2)

	if assert.NoError(t, err) && assert.NotNil(t, item) {
		assert.Equal(t, "c", item)

		expected := []string{"a", "b"}
		assertLinkedList(t, expected, l)
	}
}

func assertLinkedList(t *testing.T, expected []string, l linkedlist.SinglyLinkedList) {

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
