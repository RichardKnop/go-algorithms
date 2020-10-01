package linkedlist

import (
	"errors"
	"sync"
)

type Item interface{}

// Node a single node that composes the list
type Node struct {
	content Item
	next    *Node
}

// ItemLinkedList the linked list of Items
type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

type SinglyLinkedList interface {
	Append(t Item)                // adds an Item to the end of the linked list
	Insert(i int, t Item) error   // adds an Item at position i
	RemoveAt(i int) (Item, error) // removes a node at position i
	IndexOf(t Item) int           // returns the position of the Item t
	IsEmpty() bool                // returns true if the list is empty
	Size() int                    // returns the linked list size
	Head() *Node                  // returns the first node, so we can iterate on it
}

func New() SinglyLinkedList {
	return new(ItemLinkedList)
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Value() Item {
	return n.content
}

// Append adds an Item to the end of the linked list
func (ll *ItemLinkedList) Append(t Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	ll.size++

	node := Node{t, nil}

	if ll.head == nil {
		ll.head = &node
		return
	}

	last := ll.head
	for last.next != nil {
		last = last.next
	}
	last.next = &node
}

// Insert adds an Item at position i
func (ll *ItemLinkedList) Insert(i int, t Item) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if i < 0 || i > ll.size {
		return errors.New("index out of bounds")
	}

	ll.size++

	addNode := Node{t, nil}

	if i == 0 {
		addNode.next = ll.head
		ll.head = &addNode
		return nil
	}

	node := ll.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode

	return nil
}

// RemoveAt removes a node at position i
func (ll *ItemLinkedList) RemoveAt(i int) (Item, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if ll.head == nil {
		return nil, errors.New("cannot remove from empty list")
	}

	if i < 0 || i > ll.size {
		return nil, errors.New("index out of bounds")
	}

	ll.size--

	if i == 0 {
		remove := ll.head
		ll.head = remove.next
		return remove.content, nil
	}

	node := ll.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}

	remove := node.next
	node.next = remove.next

	return remove.content, nil
}

// IndexOf returns the position of the Item t
func (ll *ItemLinkedList) IndexOf(t Item) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	node := ll.head
	j := 0
	for {
		if node.content == t {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}

// IsEmpty returns true if the list is empty
func (ll *ItemLinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	if ll.head == nil {
		return true
	}

	return false
}

// Size returns the linked list size
func (ll *ItemLinkedList) Size() int {

	return ll.size
}

// Head returns a pointer to the first node of the list
func (ll *ItemLinkedList) Head() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.head
}
