/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package list

import (
	"fmt"

	"github.com/kevinpollet/go-datastructures/errors"
)

type elt struct {
	next  *elt
	value interface{}
}

// SinglyLinkedList is an implementation of the List ADT
type SinglyLinkedList struct {
	head *elt
	tail *elt
	size int
}

// Add appends the given value to the list
func (list *SinglyLinkedList) Add(value interface{}) {
	newElt := &elt{value: value}

	if list.size == 0 {
		list.head = newElt
	} else {
		list.tail.next = newElt
	}

	list.tail = newElt
	list.size++
}

// Clear removes all values in the list
func (list *SinglyLinkedList) Clear() {
	list.head, list.tail, list.size = nil, nil, 0
}

// Contains returns true if the list contains the given value, false otherwise
func (list *SinglyLinkedList) Contains(value interface{}) bool {
	found := false
	for it := list.head; it != nil && !found; {
		found = it.value == value
		it = it.next
	}
	return found
}

// Get returns the value at the given index in the list
func (list *SinglyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, &errors.IndexOutOfBoundsError{Index: index, Size: list.size}
	}

	it := list.head
	for i := 0; i < index; i++ {
		it = it.next
	}
	return it.value, nil
}

// Insert adds the given value at the given index in the list
func (list *SinglyLinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index > list.size {
		return &errors.IndexOutOfBoundsError{Index: index, Size: list.size}
	}

	newElt := &elt{value: value}

	if list.size == 0 {
		list.head = newElt
		list.tail = newElt

	} else if index == 0 {
		newElt.next = list.head
		list.head = newElt

	} else if index == list.size {
		list.tail.next = newElt
		list.tail = newElt

	} else {
		it := list.head
		for i := 1; i < index; i++ {
			it = it.next
		}

		newElt.next = it.next
		it.next = newElt
	}

	list.size++
	return nil
}

// IsEmpty returns true if the list is empty, false otherwise
func (list *SinglyLinkedList) IsEmpty() bool {
	return list.size == 0
}

// Join returns a string with all the concatenated and separated by the given separator
func (list *SinglyLinkedList) Join(separator string) string {
	str := ""
	if list.size == 0 {
		return str
	}
	for it := list.head; it != nil; {
		str += fmt.Sprintf("%v%s", it.value, separator)
		it = it.next
	}
	return str[:len(str)-1]
}

// Remove removes the given value from the list and returns true if the value has been removed, false otherwise
func (list *SinglyLinkedList) Remove(value interface{}) bool {
	if list.size != 0 {
		if list.head.value == value {
			list.head = list.head.next
			if list.size == 1 {
				list.tail = list.head
			}
			list.size--
			return true
		}

		for it := list.head; it.next != nil; it = it.next {
			if it.next.value == value {
				if list.tail == it.next {
					list.tail = it
				}
				it.next = it.next.next
				list.size--
				return true
			}
		}
	}

	return false
}

// Size returns the number of values in the list
func (list SinglyLinkedList) Size() int {
	return list.size
}

func (list *SinglyLinkedList) String() string {
	return fmt.Sprintf("&{head: %v, tail: %v, size: %v}", list.head, list.tail, list.size)
}
