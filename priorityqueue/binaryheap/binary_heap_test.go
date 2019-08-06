/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package binaryheap

import (
	"testing"

	"github.com/kevinpollet/go-datastructures/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func TestPriorityQueueAssertion(test *testing.T) {
	var binaryHeap interface{} = &BinaryHeap{}

	cast, ok := binaryHeap.(priorityqueue.PriorityQueue)

	assert.True(test, ok)
	assert.NotNil(test, cast)
}
