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
// Complexity: O(1)
func (dequeue *LinkedListDequeue) Clear() {
	dequeue.list.Clear()
}

// IsEmpty returns true if the dequeue is empty, false otherwise.
// Complexity: O(1)
func (dequeue *LinkedListDequeue) IsEmpty() bool {
	return dequeue.IsEmpty()
}

// OfferFirst adds the given value at the front of the dequeue.
// Complexity: O(1)
func (dequeue *LinkedListDequeue) OfferFirst(value interface{}) {
	dequeue.list.Insert(0, value)
}

// OfferLast adds the given value at the end of the dequeue.
// Complexity: O(1)
func (dequeue *LinkedListDequeue) OfferLast(value interface{}) {
	dequeue.list.Add(value)
}

// PeekFirst returns the first value in the dequeue or an error if the dequeue is empty.
// Complexity: O(1)
func (dequeue *LinkedListDequeue) PeekFirst() (interface{}, error) {
	value, err := dequeue.list.Get(0)
	if _, ok := err.(*errors.IndexOutOfBoundsError); ok {
		return nil, errors.NewNoSuchValueError("cannot peek the first value of an empty dequeue")
	}

	return value, nil
}

// PeekLast returns the last value in the dequeue or an error if the dequeue is empty.
// Complexity: O(1)
func (dequeue *LinkedListDequeue) PeekLast() (interface{}, error) {
	value, err := dequeue.list.Get(dequeue.list.Size() - 1)
	if _, ok := err.(*errors.IndexOutOfBoundsError); ok {
		return nil, errors.NewNoSuchValueError("cannot peek the last value of an empty dequeue")
	}

	return value, nil
}

// PollFirst removes and returns the first value from the dequeue or returns an error if the dequeue is empty.
// Complexity: O(1)
func (dequeue *LinkedListDequeue) PollFirst() (interface{}, error) {
	value, err := dequeue.list.Remove(0)
	if _, ok := err.(*errors.IndexOutOfBoundsError); ok {
		return nil, errors.NewNoSuchValueError("cannot poll the first value of an empty dequeue")
	}

	return value, nil
}

// PollLast removes and returns the last value from the dequeue or returns an error if the dequeue is empty.
// Complexity: O(1)
func (dequeue *LinkedListDequeue) PollLast() (interface{}, error) {
	value, err := dequeue.list.Remove(dequeue.list.Size() - 1)
	if _, ok := err.(*errors.IndexOutOfBoundsError); ok {
		return nil, errors.NewNoSuchValueError("cannot poll the last value of an empty dequeue")
	}

	return value, nil
}

// Size returns the number of values in the dequeue.
// Complexity: O(1)
func (dequeue *LinkedListDequeue) Size() int {
	return dequeue.Size()
}

// String returns a string representation of the dequeue.
// Complexity: O(n)
func (dequeue LinkedListDequeue) String() string {
	return dequeue.list.String()
}
