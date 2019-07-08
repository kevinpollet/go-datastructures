/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package list

import (
	"fmt"
)

// IndexOutOfBoundsError returned when the given index is not in [0,size[
type IndexOutOfBoundsError struct {
	index int
	size  int
}

// NewIndexOutOfBoundsError constructs and returns an instance of IndexOutOfBoundsError
func NewIndexOutOfBoundsError(index int, size int) *IndexOutOfBoundsError {
	return &IndexOutOfBoundsError{
		index: index,
		size:  size,
	}
}

func (err *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("Index out of bounds Error: index %d is not in range [0,%d[", err.index, err.size)
}
