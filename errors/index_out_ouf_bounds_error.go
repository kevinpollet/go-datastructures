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

// IndexOutOfBoundsError returned when the given index is not in [0,size[
type IndexOutOfBoundsError struct {
	Index int
	Size  int
}

func (err *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("Index out of bounds, index %d is not in range [0,%d[", err.Index, err.Size)
}
