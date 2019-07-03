/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package list

type List interface {
	Add(value interface{})
	Clear()
	Contains(value interface{}) bool
	Get(index int) (interface{}, error)
	Insert(index int, value interface{})
	IsEmpty() bool
	Join(delim string)
	Remove(value interface{}) bool
	Size() int
}
