package list

import (
	"errors"
	"fmt"
)

// List is the instance of linkedlist.
type List struct {
	head *node
}

type node struct {
	data interface{}
	next *node
}

// NewList returns a new instance of list.
func NewList() *List {
	head := new(node)
	list := &List{
		head: head,
	}
	return list
}

// HeadAdd adds a new element to the head of the linked list.
func (list *List) HeadAdd(d interface{}) {
	newNode := &node{
		data: d,
	}

	newNode.next = list.head.next
	list.head.next = newNode
}

// TailAdd adds a new element to the tail of the linked list.
func (list *List) TailAdd(d interface{}) {
	newNode := &node{
		data: d,
	}

	// find the tail node
	tail := list.head
	for tail.next != nil {
		tail = tail.next
	}

	// insert new node
	tail.next = newNode
}

// TailAdds adds batch value to the tail of list.
func (list *List) TailAdds(d ...interface{}) error {
	for _, v := range d {
		node := new(node)
		node.data = v
		list.TailAdd(node)
	}
	return nil
}

// IsEmpty returns if the list is empty.
func (list *List) IsEmpty() bool {
	return list.one() == nil
}

// Reverse will reverse the list.
func (list *List) Reverse() {
	if list.IsEmpty() {
		return
	}

	// For a node q, if we want to reverse it, we need know it's prev node p
	// and it's next node r. We should set q.next = p and store p.next in order
	// to do the same for the next node.
	p := list.one()
	q := list.one().next
	for p.next = nil; q != nil; {
		r := q.next
		q.next = p
		p = q
		q = r
	}
	list.head.next = p
}

// HasCircle returns if there is circle in the list.
func (list *List) HasCircle() bool {
	// empty list and one node list could not have circle.
	if list.IsEmpty() || list.one().next == nil {
		return false
	}

	// p step one and q step two, if p or q is nil, it means there is no circle
	// in the list, if p == q at some time, it means there has circle in the list.
	p := list.one()
	q := p.next
	for p != nil && q != nil {
		if p == q {
			return true
		}
		p = p.next
		q = q.next.next
	}
	return false
}

// Combine merges two ordered linked lists into one ordered linked list.
// TC: O(n)
func (list *List) Combine(l *List) {
	one := list.one()
	two := l.one()
	p := list.head
	for one != nil && two != nil {
		node := new(node)
		node.data = two.data
		if one.data.(int) > two.data.(int) {
			node.next = one
			p.next = node
			two = two.next
		} else {
			p = one
			one = one.next
		}
	}

	if one == nil {
		for ; two != nil; two = two.next {
			node := new(node)
			node.data = two.data
			p.next = node
			p = p.next
		}
	}
}

// DeleteFromTail will delete the nth node from tail.
func (list *List) DeleteFromTail(n int) error {
	if list.IsEmpty() {
		return errors.New("list is empty")
	}

	p := list.one()
	q := list.one()
	for i := 0; i < n; i++ {
		p = p.next
		if p == nil {
			return errors.New("list out of range")
		}
	}
	for p.next != nil {
		p = p.next
		q = q.next
	}

	q.next = q.next.next
	return nil
}

func (list *List) String() string {
	var result string
	if list.IsEmpty() {
		return ""
	}
	for p := list.one(); p != nil; p = p.next {
		result += fmt.Sprintf("%v,", p.data)
	}
	return result[:len(result)-1]
}

func (list *List) one() *node {
	return list.head.next
}

func (list *List) last() *node {
	p := list.one()
	for {
		if p.next == nil {
			return p
		}
		p = p.next
	}
}
