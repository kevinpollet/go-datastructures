/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package stack

// EmptyStackError returned when the given stack is empty
type EmptyStackError struct {
}

func (err *EmptyStackError) Error() string {
	return "Empty stack error"
}
