/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package dequeue

// Dequeue Abstract Data Types (ADT)
type Dequeue interface {
	Clear()
	IsEmpty() bool
	OfferFirst(value interface{})
	OfferLast(value interface{})
	PeekFirst() (interface{}, error)
	PeekLast() (interface{}, error)
	PollFirst() (interface{}, error)
	PollLast() (interface{}, error)
	Size() int
}
