/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package errors

import (
	"fmt"
)

// IndexOutOfBoundsError is an error returned when the given index is not in [0,size[
type IndexOutOfBoundsError struct {
	Index int
	Size  int
}

func (err *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("Index out of bounds error, index %d is not in [0,%d[", err.Index, err.Size)
}
