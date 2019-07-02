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

type LinkedList struct {
	head *cell
	tail *cell
	size int
}

func (l *LinkedList) Add(value interface{}) *LinkedList {
	cell := &cell{value: value}

	if l.IsEmpty() {
		l.head = cell
		l.tail = l.head
	} else {
		l.tail.next = cell
		l.tail = l.tail.next
	}

	l.size++
	return l
}

func (l *LinkedList) Insert(index int, value interface{}) (*LinkedList, error) {
	if index < 0 || index > l.size {
		return nil, errors.New("Index Out of Bounds Error")
	}

	if l.IsEmpty() {
		l.head = &cell{value: value}
		l.tail = l.head

	} else if index == 0 {
		l.head = &cell{next: l.head, value: value}
	} else {
		it := l.head
		for i := 0; i < index - 1; i++ {
			it = it.next
		}
		it.next = &cell{next: it.next, value: value}
	}

	l.size++
	return l, nil
}

func (l *LinkedList) Clear() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Contains(value interface{}) bool {
	for it := l.head; it != nil; {
		if it.value == value {
			return true
		}
		it = it.next
	}
	return false
}

func (l *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("Index Out of Bounds Error")
	}

	it := l.head
	for i := 0; i < index; i++ {
		it = it.next
	}

	return it.value, nil
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Join(delim string) string {
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

func (l *LinkedList) Remove(value interface{}) bool {
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
			temp.next = nil
			l.size--
			return true
		}
		it = it.next
	}

	return false
}

func (l LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) String() string {
	return fmt.Sprintf("&{ head: %v, tail: %v, size: %v }", l.head, l.tail, l.size)
}
