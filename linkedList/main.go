package main

import (
	"errors"
	"fmt"
	"strconv"
)

type elem struct {
	val  int
	prev *elem
	next *elem
}

type DblList struct {
	head  *elem
	first *elem
	last  *elem
	size  int
}

func NewList() DblList {
	return DblList{
		head:  nil,
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (l *DblList) InsertAt(idx, val int) error {
	if idx < 0 || idx > l.size-1 {
		return errors.New(fmt.Sprintf("Index %d out of bounds of size %d\n", idx, l.size))
	}

	l.head = l.first

    for range idx {
		l.head = l.head.next
	}

	var newElem elem

	if idx == 0 {
		// insert at start of list
		newElem = elem{
			val:  val,
			next: l.head,
			prev: nil,
		}

		l.head.prev = &newElem
		l.first = &newElem
	} else {
		newElem = elem{
			val:  val,
			next: l.head,
			prev: l.head.prev,
		}

		l.head.prev.next = &newElem
		l.head.prev = &newElem
	}

	l.size++
	return nil
}

func (l *DblList) Add(val int) {
	if l.size == 0 {
		newElem := elem{
			val:  val,
			next: nil,
			prev: nil,
		}

		l.first = &newElem
		l.last = &newElem
		l.size++
	} else {
		l.head = l.last

		newElem := elem{
			val:  val,
			next: nil,
			prev: l.head,
		}

		l.head.next = &newElem
		l.last = &newElem
		l.size++
	}
}

func (l *DblList) RemoveAt(idx int) error {
	if idx < 0 || idx > l.size-1 {
		return errors.New(fmt.Sprintf("Index %d out of bounds of size %d\n", idx, l.size))
	}

	l.head = l.first

	for range idx {
		l.head = l.head.next
	}

	if l.head.prev != nil {
		l.head.prev.next = l.head.next
	} else {
		l.first = l.head.next
	}

	if l.head.next != nil {
		l.head.next.prev = l.head.prev
	} else {
		l.last = l.head.prev
	}

	l.size--

	return nil
}

func (l *DblList) Clear() {
	l.head = l.first

	for range l.size {
		l.head.prev = nil

		if l.head.next != nil {
			l.head = l.head.next
			l.head.prev.next = nil
		}
	}

	l.head = nil
	l.first = nil
	l.last = nil
	l.size = 0
}

func (l *DblList) At(idx int) (int, error) {
	if idx < 0 || idx > l.size-1 {
		return -1, errors.New(fmt.Sprintf("Index %d out of bounds of size %d\n", idx, l.size))
	}

	l.head = l.first

	for range idx {
		l.head = l.head.next
	}

	return l.head.val, nil
}

func (l *DblList) ToSlice() []int {
	out := make([]int, l.size)

	l.head = l.first

    for i := range l.size {
		out[i] = l.head.val
		l.head = l.head.next
	}

	l.Clear()
	return out
}

func (l *DblList) Reverse() {
	l.head = l.first
	l.last = l.head

	for range l.size {
		l.head.next, l.head.prev = l.head.prev, l.head.next

		if l.head.prev != nil {
			l.head = l.head.prev
		}
	}

	l.first = l.head
}

func (l *DblList) String() string {
	out := "["

	l.head = l.first

	for i := range l.size {
		out += strconv.Itoa(l.head.val)

		if i != l.size-1 {
			out += ", "
		}

		l.head = l.head.next
	}

	out += "]\n"

	return out
}

func main() {
	myList := NewList()

	myList.Add(0)
	// insert at the start
	if err := myList.InsertAt(0, 1); err != nil {
		panic(err)
	}
	fmt.Println(myList.String())

	// add elements
	myList.Add(2)
	myList.Add(3)
	myList.Add(4)
	myList.Add(5)
	fmt.Println(myList.String())

	// remove an element
	if err := myList.RemoveAt(0); err != nil {
		panic(err)
	}
	fmt.Println(myList.String())

	// insert in the middle
	if err := myList.InsertAt(1, 1); err != nil {
		panic(err)
	}
	fmt.Println(myList.String())

	// reverse the list
	myList.Reverse()
	fmt.Println(myList.String())

	// get an element
	val, err := myList.At(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("val at %d: %d\n", 1, val)

	// convert to slice
	testSlice := myList.ToSlice()
	for i, e := range testSlice {
		fmt.Printf("slice val at %d: %d\n", i, e)
	}

	// clear the list
	myList.Clear()
	fmt.Println(myList.String())
}
