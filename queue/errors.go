/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package queue

// EmptyQueueError is returned when the queue is empty
type EmptyQueueError struct {
}

// NewEmptyQueueError constructs and returns an instance of EmptyQueueError
func NewEmptyQueueError() *EmptyQueueError {
	return &EmptyQueueError{}
}

func (err *EmptyQueueError) Error() string {
	return "Empty Queue Error"
}
