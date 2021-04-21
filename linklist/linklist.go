package linklist

import (
	"fmt"
)

// ListNode ...
type ListNode struct {
	data interface{}
	next *ListNode
}

// LinkedList ...
type LinkedList struct {
	head *ListNode
	size int
}

var ErrListEmpty = fmt.Errorf("list is empty")
var ErrIndexOutOfBound = fmt.Errorf("index out of bound")

// Display print all the link list.
func (ll *LinkedList) Display() error {
	if ll.head == nil {
		return ErrListEmpty
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v, ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

// Length return length of link list.
func (ll *LinkedList) Length() int {
	return ll.size
}

// InsertBegin insert element at the begining.
func (ll *LinkedList) InsertBegin(data interface{}) {
	node := &ListNode{
		data: data,
	}
	// case empty link list
	if ll.head == nil {
		ll.head = node
		// case not empty
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
}

// InsertEnd insert element at the end.
func (ll *LinkedList) InsertEnd(data interface{}) {
	node := &ListNode{
		data: data,
	}

	// case empty link list
	if ll.head == nil {
		ll.head = node
		// case not empty
	} else {
		current := ll.head
		for ; current.next != nil; current = current.next {
		}
		current.next = node
	}
	ll.size++
}

func (ll *LinkedList) Insert(pos int, data interface{}) error {
	if pos < 0 || pos > ll.size {
		return ErrIndexOutOfBound
	}
	if pos == 0 {
		ll.InsertBegin(data)
		return nil
	}
	node := &ListNode{
		data: data,
	}
	var prev, current *ListNode
	current = ll.head
	for i := 0; i < pos; i++ {
		prev = current
		current = current.next
	}
	prev.next = node
	node.next = current
	ll.size++
	return nil
}
