/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedlistdequeue

import (
	"testing"

	"github.com/kevinpollet/go-datastructures/dequeue"
	"github.com/stretchr/testify/assert"
)

func TestDequeueAssertion(t *testing.T) {
	var linkedListDequeue interface{} = &LinkedListDequeue{}

	value, ok := linkedListDequeue.(dequeue.Dequeue)

	assert.True(t, ok)
	assert.NotNil(t, value)
}
