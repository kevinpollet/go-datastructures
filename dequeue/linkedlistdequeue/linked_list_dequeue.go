/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedlistdequeue

import (
	"github.com/kevinpollet/go-datastructures/errors"
	"github.com/kevinpollet/go-datastructures/list/doublylinkedlist"
)

// LinkedListDequeue implements the Dequeue ADT with a DoublyLinkedList.
type LinkedListDequeue struct {
	list doublylinkedlist.DoublyLinkedList
}

// Clear removes all values from the dequeue.
func (dequeue *LinkedListDequeue) Clear() {
	dequeue.list.Clear()
}

// DequeueFirst removes and returns the first value from the dequeue or returns an error if the dequeue is empty.
func (dequeue *LinkedListDequeue) DequeueFirst() (interface{}, error) {
	if dequeue.list.IsEmpty() {
		return nil, errors.NewNoSuchElementError("cannot dequeue the first value of an empty dequeue")
	}

	value, _ := dequeue.list.Get(0)
	dequeue.list.Remove(0)
	return value, nil
}

// DequeueLast removes and returns the last value from the dequeue or returns an error if the dequeue is empty.
func (dequeue *LinkedListDequeue) DequeueLast() (interface{}, error) {
	if dequeue.list.IsEmpty() {
		return nil, errors.NewNoSuchElementError("cannot dequeue the last value of an empty dequeue")
	}

	index := dequeue.list.Size() - 1
	value, _ := dequeue.list.Get(index)
	dequeue.list.Remove(index)
	return value, nil
}

// Append adds the given value at the end of the dequeue.
func (dequeue *LinkedListDequeue) Append(value interface{}) {
	dequeue.list.Add(value)
}

// IsEmpty returns true if the dequeue is empty, false otherwise.
func (dequeue *LinkedListDequeue) IsEmpty() bool {
	return dequeue.IsEmpty()
}

// PeekFirst returns the first value in the dequeue or an error if the dequeue is empty.
func (dequeue *LinkedListDequeue) PeekFirst() (interface{}, error) {
	if dequeue.IsEmpty() {
		return nil, errors.NewNoSuchElementError("cannot peek the first value of an empty dequeue")
	}
	value, _ := dequeue.list.Get(0)
	return value, nil
}

// PeekLast returns the last value in the dequeue or an error if the dequeue is empty.
func (dequeue *LinkedListDequeue) PeekLast() (interface{}, error) {
	if dequeue.IsEmpty() {
		return nil, errors.NewNoSuchElementError("cannot peek the last value of an empty dequeue")
	}
	value, _ := dequeue.list.Get(dequeue.list.Size() - 1)
	return value, nil
}

// Prepend adds the given value at the front of the dequeue.
func (dequeue *LinkedListDequeue) Prepend(value interface{}) {
	dequeue.list.Insert(0, value)
}

// Size returns the number of values in the dequeue.
func (dequeue *LinkedListDequeue) Size() int {
	return dequeue.Size()
}

func (dequeue LinkedListDequeue) String() string {
	return dequeue.list.String()
}
