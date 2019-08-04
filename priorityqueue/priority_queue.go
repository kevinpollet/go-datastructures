/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package priorityqueue

// PriorityQueue Abstract Data Type (ADT)
type PriorityQueue interface {
	Clear()
	Insert(value interface{}, priority int)
	IsEmpty()
	Peek() (interface{}, error)
	Poll() (interface{}, error)
	Size()
}
