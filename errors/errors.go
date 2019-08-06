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

// NoSuchElementError is returned to indicate that there are no values in the ADT implementation
type NoSuchElementError struct {
	detail string
}

// NewNoSuchElementError constructs and returns an instance of NoSuchElementError
func NewNoSuchElementError(detail string) *NoSuchElementError {
	return &NoSuchElementError{detail: detail}
}

func (err *NoSuchElementError) Error() string {
	if err.detail == "" {
		return fmt.Sprintf("NoSuchElementError")
	}
	return fmt.Sprintf("NoSuchElementError: %s", err.detail)
}

// IndexOutOfBoundsError is returned that the given index is out of bounds
type IndexOutOfBoundsError struct {
	index int
	size  int
}

// NewIndexOutOfBoundsError constructs and returns an instance of IndexOutOfBoundsError
func NewIndexOutOfBoundsError(index, size int) *IndexOutOfBoundsError {
	return &IndexOutOfBoundsError{
		index: index,
		size:  size,
	}
}

func (err *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("IndexOutOfBoundsError: index %d, size %d", err.index, err.size)
}
