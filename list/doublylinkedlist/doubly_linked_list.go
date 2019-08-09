/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package doublylinkedlist

import (
	"fmt"

	"github.com/kevinpollet/go-datastructures/errors"
)

type elt struct {
	next  *elt
	prev  *elt
	value interface{}
}

// DoublyLinkedList implements the List ADT.
type DoublyLinkedList struct {
	head *elt
	tail *elt
	size int
}

// Add appends the given value to the list.
func (list *DoublyLinkedList) Add(value interface{}) {
	if list.IsEmpty() {
		list.head = &elt{value: value}
		list.tail = list.head
	} else {
		newElement := &elt{prev: list.tail, value: value}
		list.tail.next = newElement
		list.tail = newElement
	}
	list.size++
}

// Clear removes all values from the list.
func (list *DoublyLinkedList) Clear() {
	list.head, list.tail, list.size = nil, nil, 0
}

// IndexOf returns the index of the first occurrence of the given value in the list, -1 if the list does not contain the value.
func (list *DoublyLinkedList) IndexOf(value interface{}) int {
	index := -1
	for i, it := 0, list.head; index == -1 && it != nil; i, it = i+1, it.next {
		if it.value == value {
			index = i
		}
	}
	return index
}

// Insert inserts the given value at the given index or returns an error if the given index is out of bounds.
func (list *DoublyLinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index > list.Size() {
		return errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	if index == list.Size() {
		list.Add(value)

	} else if index == 0 {
		newElement := &elt{next: list.head, value: value}
		list.head.prev = newElement
		list.head = newElement
		list.size++

	} else {
		it := list.head.next
		for i := 1; i < index; i++ {
			it = it.next
		}

		newElement := &elt{next: it, prev: it.prev, value: value}
		newElement.prev.next = newElement
		if it != list.tail {
			newElement.next.prev = newElement
		}
		list.size++
	}
	return nil
}

// IsEmpty return true if the list is empty, false otherwise.
func (list *DoublyLinkedList) IsEmpty() bool {
	return list.size == 0
}

// Get returns the value at the given index or an error if the given index is out of bounds.
func (list *DoublyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.Size() {
		return nil, errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	it := list.head
	for i := 1; i <= index; i++ {
		it = it.next
	}
	return it.value, nil
}

// Remove removes the value at the given index from the list or returns an error if the given index is out of bounds.
func (list *DoublyLinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= list.Size() {
		return nil, errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	var removedValue interface{}

	it := list.head
	for i := 0; i < index; i++ {
		it = it.next
	}

	if it == list.head {
		removedValue = list.head.value
		list.head = list.head.next

		if list.tail == it {
			list.tail = list.head
		} else {
			list.head.prev = nil
		}

	} else if it == list.tail {
		removedValue = list.tail.value
		list.tail = list.tail.prev
		list.tail.next = nil

	} else {
		removedValue = it.value
		it.prev.next = it.next
		it.next.prev = it.prev
	}

	list.size--

	return removedValue, nil
}

// Size return the number of value in the list.
func (list *DoublyLinkedList) Size() int {
	return list.size
}

func (list DoublyLinkedList) String() string {
	str := "["
	for it := list.head; it != nil; it = it.next {
		str += fmt.Sprintf("%v", it.value)
		if it != list.tail {
			str += ","
		}
	}
	str += "]"

	return str
}
