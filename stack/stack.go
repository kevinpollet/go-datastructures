/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package stack

// Stack ADS
type Stack interface {
	Clear()
	IsEmpty() bool
	Peek() interface{}
	Pop() (interface{}, error)
	Push(value interface{})
	Size() int
}
