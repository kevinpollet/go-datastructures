/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedlistqueue

import (
	"testing"

	"github.com/kevinpollet/go-datastructures/list/doublylinkedlist"
	queuePkg "github.com/kevinpollet/go-datastructures/queue"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name   string
	values []int
}{
	{"EmptyQueue", []int{}},
	{"QueueWithOneValue", []int{1}},
	{"QueueWithTwoValues", []int{1, 2}},
}

func newLinkedListQueue(values ...int) LinkedListQueue {
	list := doublylinkedlist.DoublyLinkedList{}
	for _, value := range values {
		list.Add(value)
	}
	return LinkedListQueue{list: list}
}

func TestQueueAssertion(test *testing.T) {
	var queue interface{} = &LinkedListQueue{}

	value, ok := queue.(queuePkg.Queue)

	assert.True(test, ok)
	assert.NotNil(test, value)
}

func TestClear(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			queue := newLinkedListQueue()
			queue.Clear()

			assert.Equal(t, 0, queue.list.Size())
		})
	}
}

func TestDequeue(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			queue := newLinkedListQueue(testCase.values...)
			value, err := queue.Dequeue()

			switch {
			case len(testCase.values) == 0:
				assert.Nil(t, value)
				assert.Error(t, err)
			default:
				assert.Nil(t, err)
				assert.Equal(t, testCase.values[0], value)
			}
		})
	}
}

func TestEnqueue(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			queue := newLinkedListQueue()
			for _, value := range testCase.values {
				queue.Enqueue(value)
			}

			assert.Equal(t, len(testCase.values), queue.list.Size())
			for idx, value := range testCase.values {
				listValue, _ := queue.list.Get(idx)
				assert.Equal(t, value, listValue)
			}
		})
	}
}

func TestIsEmpty(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			queue := newLinkedListQueue(testCase.values...)

			assert.Equal(t, len(testCase.values) == 0, queue.IsEmpty())
		})
	}
}

func TestPeek(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			queue := newLinkedListQueue(testCase.values...)
			value, err := queue.Peek()

			switch {
			case len(testCase.values) == 0:
				assert.Nil(t, value)
				assert.Error(t, err)
			default:
				assert.Equal(t, testCase.values[0], value)
			}
		})
	}
}

func TestSize(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			queue := newLinkedListQueue(testCase.values...)

			assert.Equal(t, len(testCase.values), queue.Size())
		})
	}
}
