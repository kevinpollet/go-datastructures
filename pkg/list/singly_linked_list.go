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

type cell struct {
	next  *cell
	value interface{}
}

type SinglyLinkedList struct {
	head *cell
	tail *cell
	size int
}

func (l *SinglyLinkedList) Add(value interface{}) {
	cell := &cell{value: value}

	if l.IsEmpty() {
		l.head = cell
		l.tail = l.head
	} else {
		l.tail.next = cell
		l.tail = l.tail.next
	}

	l.size++
}

func (l *SinglyLinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index > l.size {
		return errors.New("Index Out of Bounds Error")
	}

	if l.IsEmpty() {
		l.head = &cell{value: value}
		l.tail = l.head

	} else if index == 0 {
		l.head = &cell{next: l.head, value: value}
	} else {
		it := l.head
		for i := 0; i < index-1; i++ {
			it = it.next
		}
		it.next = &cell{next: it.next, value: value}
	}

	l.size++
	return nil
}

func (l *SinglyLinkedList) Clear() {
	*l = SinglyLinkedList{}
}

func (l *SinglyLinkedList) Contains(value interface{}) bool {
	for it := l.head; it != nil; {
		if it.value == value {
			return true
		}
		it = it.next
	}
	return false
}

func (l *SinglyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("Index Out of Bounds Error")
	}

	it := l.head
	for i := 0; i < index; i++ {
		it = it.next
	}

	return it.value, nil
}

func (l *SinglyLinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *SinglyLinkedList) Join(delim string) string {
	if l.IsEmpty() {
		return ""
	}

	var str string
	for it := l.head; it != nil; {
		str += fmt.Sprintf("%v%s", it.value, delim)
		it = it.next
	}

	return str[:len(str)-1]
}

func (l *SinglyLinkedList) Remove(value interface{}) bool {
	if l.IsEmpty() {
		return false
	}

	if l.head.value == value {
		temp := l.head
		l.head = temp.next
		if l.tail == temp {
			l.tail = l.head
		}
		l.size--
		return true
	}

	for it := l.head; it.next != nil; {
		if it.next.value == value {
			temp := it.next
			it.next = temp.next
			l.size--
			if l.tail == temp {
				l.tail = it
			}
			temp.next = nil
			return true
		}
		it = it.next
	}

	return false
}

func (l SinglyLinkedList) Size() int {
	return l.size
}

func (l *SinglyLinkedList) String() string {
	return fmt.Sprintf("&{ head: %v, tail: %v, size: %v }", l.head, l.tail, l.size)
}
