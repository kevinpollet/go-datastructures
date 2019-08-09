/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package errors

import "fmt"

// NoSuchValueError is returned to indicate that the data structure has no values.
type NoSuchValueError struct {
	detail string
}

// NewNoSuchValueError constructs and returns an instance of NoSuchValueError.
func NewNoSuchValueError(detail string) *NoSuchValueError {
	return &NoSuchValueError{detail: detail}
}

func (err *NoSuchValueError) Error() string {
	return fmt.Sprintf("NoSuchValueError: %s", err.detail)
}

// IndexOutOfBoundsError is returned when the given index is not within bounds [0, max[.
type IndexOutOfBoundsError struct {
	index int
	max   int
}

// NewIndexOutOfBoundsError constructs and returns an instance of IndexOutOfBoundsError.
func NewIndexOutOfBoundsError(index, max int) *IndexOutOfBoundsError {
	return &IndexOutOfBoundsError{
		index: index,
		max:   max,
	}
}

func (err *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("IndexOutOfBoundsError: index %v, bounds [0,%v[", err.index, err.max)
}
