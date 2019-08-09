/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package priorityqueue

// PriorityQueue Abstract Data Type (ADT).
type PriorityQueue interface {
	Clear()
	IsEmpty() bool
	Offer(value interface{}, priority int)
	Peek() (interface{}, error)
	Poll() (interface{}, error)
	Size() int
}
