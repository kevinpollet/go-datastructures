/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package doublylinkedlist

import (
	"fmt"

	listp "github.com/kevinpollet/go-datastructures/list"
)

type elt struct {
	next  *elt
	prev  *elt
	value interface{}
}

// DoublyLinkedList implementation of list Abstract Data Structure
type DoublyLinkedList struct {
	head *elt
	tail *elt
	size int
}

// Add append the given value to the list
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

// Clear removes all values
func (list *DoublyLinkedList) Clear() {
	list.head, list.tail, list.size = nil, nil, 0
}

// Contains returns true if the list contains the given value
func (list *DoublyLinkedList) Contains(value interface{}) bool {
	found := false
	for it := list.head; !found && it != nil; it = it.next {
		found = it.value == value
	}
	return found
}

// IsEmpty return true if the list is empty, false otherwise
func (list *DoublyLinkedList) IsEmpty() bool {
	return list.size == 0
}

// Get returns the value at the given index.
// If the given index is not in the range [0, list.Size()[ an error is returned.
func (list *DoublyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.Size() {
		return nil, &listp.IndexOutOfBoundsError{Index: index, Size: list.Size()}
	}

	it := list.head
	for i := 1; i <= index; i++ {
		it = it.next
	}
	return it.value, nil
}

// Insert inserts the given value at the given index.
// If the given index is not in the range [0, list.Size()] an error is returned.
func (list *DoublyLinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index > list.Size() {
		return &listp.IndexOutOfBoundsError{Index: index, Size: list.Size()}
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

// Join returns a string containing all values separated by the given separator
func (list *DoublyLinkedList) Join(separator string) string {
	str := ""

	for it := list.head; it != nil; it = it.next {
		str += fmt.Sprintf("%v", it.value)
		if it.next != nil {
			str += separator
		}
	}

	return str
}

// Remove removes the given value and returns true if value exists, false otherwise
func (list *DoublyLinkedList) Remove(value interface{}) bool {
	found := false
	if !list.IsEmpty() {
		it := list.head
		for it != nil && it.value != value {
			it = it.next
		}

		found = it != nil
		if found {
			if it != list.head {
				it.prev.next = it.next
			} else {
				list.head = it.next
			}
			if it != list.tail {
				it.next.prev = it.prev
			} else {
				list.tail = it.prev
			}
			list.size--
		}

	}
	return found
}

// Size return the number of elements in the list
func (list *DoublyLinkedList) Size() int {
	return list.size
}

func (list DoublyLinkedList) String() string {
	return fmt.Sprintf("[%s]", list.Join(","))
}
