/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package list

// List Abstract Data Type (ADT)
type List interface {
	Add(value interface{})
	AddAtIndex(index int, value interface{}) error
	Clear()
	Contains(value interface{}) bool
	Get(index int) (interface{}, error)
	IsEmpty() bool
	Join(separator string) string
	Remove(value interface{}) bool
	RemoveAtIndex(index int) error
	Size() int
}
