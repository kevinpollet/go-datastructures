/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package list

// List Abstract Data Type (ADT).
type List interface {
	Add(value interface{})
	Clear()
	Contains(value interface{}) bool
	Get(index int) (interface{}, error)
	Insert(index int, value interface{}) error
	IsEmpty() bool
	Remove(value interface{}) bool
	RemoveAtIndex(index int) error
	Size() int
}
