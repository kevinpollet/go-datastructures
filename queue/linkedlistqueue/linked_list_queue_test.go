/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedlistqueue

import (
	"testing"

	queuePackage "github.com/kevinpollet/go-datastructures/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueueImplementation(test *testing.T) {
	var queue interface{} = &LinkedListQueue{}

	cast, ok := queue.(queuePackage.Queue)

	assert.True(test, ok)
	assert.NotNil(test, cast)
}
