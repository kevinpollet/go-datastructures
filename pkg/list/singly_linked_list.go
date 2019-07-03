/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package list

import (
	"errors"
	"fmt"
)

type elt struct {
	next  *elt
	value interface{}
}

type SinglyLinkedList struct {
	head *elt
	tail *elt
	size int
}

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

func (list *SinglyLinkedList) Clear() {
	list.head, list.tail, list.size = nil, nil, 0
}

func (list *SinglyLinkedList) Contains(value interface{}) bool {
	found := false
	for it := list.head; it != nil && !found; {
		found = it.value == value
		it = it.next
	}
	return found
}

func (list *SinglyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("Index Out of Bounds Error")
	}

	it := list.head
	for i := 0; i < index; i++ {
		it = it.next
	}
	return it.value, nil
}

func (list *SinglyLinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index > list.size {
		return errors.New("Index Out of Bounds Error")
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

func (list *SinglyLinkedList) IsEmpty() bool {
	return list.size == 0
}

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

func (list SinglyLinkedList) Size() int {
	return list.size
}

func (list *SinglyLinkedList) String() string {
	return fmt.Sprintf("&{head: %v, tail: %v, size: %v}", list.head, list.tail, list.size)
}
