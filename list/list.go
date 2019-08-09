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
	Get(index int) (interface{}, error)
	IndexOf(value interface{}) int
	Insert(index int, value interface{}) error
	IsEmpty() bool
	Remove(index int) error
	Size() int
}
