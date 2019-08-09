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
// Complexity: O(1).
func (list *DoublyLinkedList) Add(value interface{}) {
	eltToAdd := elt{value: value}

	if list.IsEmpty() {
		list.head = &eltToAdd
		list.tail = list.head
	} else {
		eltToAdd.prev = list.tail
		list.tail.next = &eltToAdd
		list.tail = &eltToAdd
	}

	list.size++
}

// Clear removes all values from the list.
// Complexity: O(1).
func (list *DoublyLinkedList) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Get returns the value at the given index or an error if the given index is out of bounds.
// Complexity: O(1) to retrieve head and tail values, O(n) otherwise.
func (list *DoublyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.Size() {
		return nil, errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	if index == list.Size()-1 {
		return list.tail.value, nil
	}

	it := list.head
	for i := 1; it != nil && i <= index; i++ {
		it = it.next
	}

	return it.value, nil
}

// IndexOf returns the index of the first occurrence of the given value in the list, -1 if the list does not contain the value.
// Complexity: O(n).
func (list *DoublyLinkedList) IndexOf(value interface{}) int {
	index := -1
	for i, it := 0, list.head; it != nil; i, it = i+1, it.next {
		if it.value == value {
			index = i
			break
		}
	}
	return index
}

// Insert inserts the given value at the given index or returns an error if the given index is out of bounds.
// Complexity: O(1) if the value is inserted at the head or tail of the list, O(n) otherwise.
func (list *DoublyLinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index > list.Size() {
		return errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	if list.IsEmpty() || index == list.Size() {
		list.Add(value)

	} else {
		eltToInsert := elt{value: value}

		if index == 0 {
			eltToInsert.next = list.head
			list.head.prev = &eltToInsert
			list.head = &eltToInsert

		} else {
			it := list.head
			for i := 1; i <= index; i++ {
				it = it.next
			}

			eltToInsert.next = it
			eltToInsert.prev = it.prev
			it.prev.next = &eltToInsert
			it.prev = &eltToInsert
		}

		list.size++
	}

	return nil
}

// IsEmpty return true if the list is empty, false otherwise.
// Complexity: O(1).
func (list *DoublyLinkedList) IsEmpty() bool {
	return list.size == 0
}

// Remove removes the value at the given index from the list or returns an error if the given index is out of bounds.
// Complexity: O(1) if the value is removed at the head or tail of the list, O(n) otherwise.
func (list *DoublyLinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= list.Size() {
		return nil, errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	var removedValue interface{}

	if list.Size() == 1 {
		removedValue = list.head.value
		list.Clear()

	} else {
		if index == 0 {
			removedValue = list.head.value
			list.head = list.head.next
			list.head.prev = nil

		} else if index == list.Size()-1 {
			removedValue = list.tail.value
			list.tail = list.tail.prev
			list.tail.next = nil

		} else {
			it := list.head.next
			for i := 1; i < index; i++ {
				it = it.next
			}

			removedValue = it.value
			it.prev.next = it.next
			it.next.prev = it.prev
		}

		list.size--
	}

	return removedValue, nil
}

// Size return the number of value in the list.
// Complexity: O(1).
func (list *DoublyLinkedList) Size() int {
	return list.size
}

// String returns a string representation of the list.
// Complexity: O(n).
func (list DoublyLinkedList) String() string {
	str := "["
	for it := list.head; it != nil; it = it.next {
		str += fmt.Sprintf("%v", it.value)
		if it != list.tail {
			str += ","
		}
	}

	return str + "]"
}
