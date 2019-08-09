/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package singlylinkedlist

import (
	"fmt"

	"github.com/kevinpollet/go-datastructures/errors"
)

type elt struct {
	next  *elt
	value interface{}
}

// SinglyLinkedList implements the List ADT.
type SinglyLinkedList struct {
	head *elt
	tail *elt
	size int
}

// Add appends the given value to the list.
func (list *SinglyLinkedList) Add(value interface{}) {
	if list.IsEmpty() {
		list.head = &elt{value: value}
		list.tail = list.head
	} else {
		list.tail.next = &elt{value: value}
		list.tail = list.tail.next
	}
	list.size++
}

// AddAtIndex adds the given value at the given index in the list or returns an error if the given index is out of bounds.
func (list *SinglyLinkedList) AddAtIndex(index int, value interface{}) error {
	if index < 0 || index > list.Size() {
		return errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	newElement := &elt{value: value}

	if index == list.Size() {
		list.Add(value)

	} else {
		it := list.head
		for i := 1; i < index; i++ {
			it = it.next
		}

		if index == 0 && it == list.head {
			newElement.next = list.head
			list.head = newElement
		} else {
			newElement.next = it.next
			it.next = newElement
		}

		list.size++
	}
	return nil
}

// Clear removes all values from the list.
func (list *SinglyLinkedList) Clear() {
	list.head, list.tail, list.size = nil, nil, 0
}

// Contains returns true if the list contains the given value, false otherwise.
func (list *SinglyLinkedList) Contains(value interface{}) bool {
	found := false
	for it := list.head; !found && it != nil; it = it.next {
		found = it.value == value
	}
	return found
}

// Get returns the value at the given index in the list or an error if the given index is out of bounds.
func (list *SinglyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.Size() {
		return nil, errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	it := list.head
	for i := 1; i <= index; i++ {
		it = it.next
	}
	return it.value, nil
}

// IsEmpty returns true if the list is empty, false otherwise.
func (list *SinglyLinkedList) IsEmpty() bool {
	return list.size == 0
}

// Remove removes the given value in the list and returns true if the value has been removed, false otherwise.
func (list *SinglyLinkedList) Remove(value interface{}) bool {
	found := false

	if !list.IsEmpty() {
		it := list.head
		if it.value == value {
			list.head = list.head.next
			if list.tail == it {
				list.tail = list.head
			}
			found = true
			list.size--
		} else {
			for it.next != nil && it.next.value != value {
				it = it.next
			}
			found = it.next != nil
			if found {
				if it.next == list.tail {
					list.tail = it
				}
				it.next = it.next.next
				list.size--
			}
		}
	}

	return found
}

// RemoveAtIndex removes the value at the given index from the list or returns an error if the given index is out of bounds.
func (list *SinglyLinkedList) RemoveAtIndex(index int) error {
	if index < 0 || index >= list.Size() {
		return errors.NewIndexOutOfBoundsError(index, list.Size())
	}

	it := list.head
	for i := 1; i < index; i++ {
		it = it.next
	}

	if list.tail == it.next {
		list.tail = it

	} else if list.tail == list.head {
		list.tail = list.head.next
	}

	if index == 0 {
		list.head = list.head.next
	} else {
		it.next = it.next.next
	}

	list.size--

	return nil
}

// Size returns the number of values in the list.
func (list SinglyLinkedList) Size() int {
	return list.size
}

func (list SinglyLinkedList) String() string {
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
