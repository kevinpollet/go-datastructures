/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package queue

// Queue Abstract Data Type (ADT).
type Queue interface {
	Clear()
	Dequeue() (interface{}, error)
	Enqueue(value interface{})
	IsEmpty() bool
	Peek() (interface{}, error)
	Size() int
}
