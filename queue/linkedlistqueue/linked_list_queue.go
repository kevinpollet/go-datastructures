/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedlistqueue

import (
	"github.com/kevinpollet/go-datastructures/list/doublylinkedlist"
	queuePackage "github.com/kevinpollet/go-datastructures/queue"
)

// LinkedListQueue implements the Queue ADT with a doubly linked list
type LinkedListQueue struct {
	list doublylinkedlist.DoublyLinkedList
}

// Clear removes all elements in the queue
func (queue *LinkedListQueue) Clear() {
	queue.list.Clear()
}

// Dequeue removes and returns the first element in the queue or returns an error if the queue is empty
func (queue *LinkedListQueue) Dequeue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, queuePackage.NewEmptyQueueError()
	}

	elt, _ := queue.list.Get(0)
	queue.list.RemoveAtIndex(0)
	return elt, nil
}

// Enqueue adds the given element to the queue
func (queue *LinkedListQueue) Enqueue(value interface{}) {
	queue.list.Add(value)
}

// IsEmpty returns true if the queue is empty, false otherwise
func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.list.IsEmpty()
}

// Peek returns the first element in the queue or returns an error if the queue is empty
func (queue *LinkedListQueue) Peek() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, queuePackage.NewEmptyQueueError()
	}
	return queue.list.Get(0)
}

// Size returns the number of elements in the queue
func (queue *LinkedListQueue) Size() int {
	return queue.list.Size()
}

func (queue LinkedListQueue) String() string {
	return queue.list.String()
}
