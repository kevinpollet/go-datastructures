/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package priorityqueue

// EmptyPriorityQueueError is returned when the priority queue is empty
type EmptyPriorityQueueError struct {
}

// NewEmptyPriorityQueueError constructs and returns an instance of EmptyPriorityQueueError
func NewEmptyPriorityQueueError() *EmptyPriorityQueueError {
	return &EmptyPriorityQueueError{}
}

func (err *EmptyPriorityQueueError) Error() string {
	return "Empty Priority Queue Error"
}
