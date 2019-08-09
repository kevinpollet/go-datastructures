/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedlistqueue

import (
	"github.com/kevinpollet/go-datastructures/errors"
	"github.com/kevinpollet/go-datastructures/list/doublylinkedlist"
)

// LinkedListQueue implements the Queue ADT with a DoublyLinkedList.
type LinkedListQueue struct {
	list doublylinkedlist.DoublyLinkedList
}

// Clear removes all values from the queue.
// Complexity: O(1)
func (queue *LinkedListQueue) Clear() {
	queue.list.Clear()
}

// IsEmpty returns true if the queue is empty, false otherwise.
// Complexity: O(1)
func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.list.IsEmpty()
}

// Offer adds the given value to the queue.
// Complexity: O(1)
func (queue *LinkedListQueue) Offer(value interface{}) {
	queue.list.Add(value)
}

// Peek returns the first value in the queue or returns an error if the queue is empty.
// Complexity: O(1)
func (queue *LinkedListQueue) Peek() (interface{}, error) {
	value, err := queue.list.Get(0)
	if _, ok := err.(*errors.IndexOutOfBoundsError); ok {
		return nil, errors.NewNoSuchValueError("cannot peek a value from an empty queue")
	}

	return value, nil
}

// Poll removes and returns the first value in the queue or returns an error if the queue is empty.
// Complexity: O(1)
func (queue *LinkedListQueue) Poll() (interface{}, error) {
	value, err := queue.list.Remove(0)
	if _, ok := err.(*errors.IndexOutOfBoundsError); ok {
		return nil, errors.NewNoSuchValueError("cannot poll a value from an empty queue")
	}

	return value, nil
}

// Size returns the number of values in the queue.
// Complexity: O(1)
func (queue *LinkedListQueue) Size() int {
	return queue.list.Size()
}

// String returns a string representation of this queue.
// Complexity: O(1)
func (queue LinkedListQueue) String() string {
	return queue.list.String()
}
