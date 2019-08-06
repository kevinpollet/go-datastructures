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

// LinkedListQueue implements the Queue ADT with a doubly linked list
type LinkedListQueue struct {
	list doublylinkedlist.DoublyLinkedList
}

// Clear removes all values in the queue
func (queue *LinkedListQueue) Clear() {
	queue.list.Clear()
}

// Dequeue removes and returns the first value in the queue or returns an error if the queue is empty
func (queue *LinkedListQueue) Dequeue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.NewNoSuchElementError("cannot dequeue a value from an empty queue")
	}

	elt, _ := queue.list.Get(0)
	queue.list.RemoveAtIndex(0)
	return elt, nil
}

// Enqueue adds the given value to the queue
func (queue *LinkedListQueue) Enqueue(value interface{}) {
	queue.list.Add(value)
}

// IsEmpty returns true if the queue is empty, false otherwise
func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.list.IsEmpty()
}

// Peek returns the first value in the queue or returns an error if the queue is empty
func (queue *LinkedListQueue) Peek() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.NewNoSuchElementError("cannot peek a value from an empty queue")
	}
	return queue.list.Get(0)
}

// Size returns the number of values in the queue
func (queue *LinkedListQueue) Size() int {
	return queue.list.Size()
}

func (queue LinkedListQueue) String() string {
	return queue.list.String()
}
