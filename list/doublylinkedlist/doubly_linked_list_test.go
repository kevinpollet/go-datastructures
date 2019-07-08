/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// []
	list := DoublyLinkedList{}
	list.Add(1)

	assert.Equal(t, 1, list.size)
	assert.Equal(t, 1, list.head.value)
	assert.Equal(t, list.tail, list.head)
	assert.Nil(t, list.tail.next)

	// [1]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)

	assert.Equal(t, 2, list.size)
	assert.Equal(t, 1, list.head.value)
	assert.Equal(t, list.tail, list.head.next)
	assert.Equal(t, 2, list.tail.value)
	assert.Nil(t, list.tail.next)
}

func TestClear(t *testing.T) {
	// []
	list := DoublyLinkedList{}
	list.Clear()

	assert.Equal(t, 0, list.size)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

	// [1]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Clear()

	assert.Equal(t, 0, list.size)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

	// [1, 2]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)
	list.Clear()

	assert.Equal(t, 0, list.size)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func TestContains(t *testing.T) {
	// []
	assert.False(t, (&DoublyLinkedList{}).Contains("foo"))

	// ["foo"]
	list := DoublyLinkedList{}
	list.Add("foo")

	assert.True(t, list.Contains("foo"))
	assert.False(t, list.Contains("bar"))
}

func TestGet(t *testing.T) {
	// index < 0
	_, err := (&DoublyLinkedList{}).Get(-1)

	assert.Error(t, err)

	// index >= size
	_, err = (&DoublyLinkedList{size: 0}).Get(0)

	assert.Error(t, err, "Index Out of Bounds Error")

	// [1]
	list := DoublyLinkedList{}
	list.Add(1)
	res, err := list.Get(0)

	assert.Equal(t, 1, res)
	assert.Nil(t, err)

}

func TestInsert(t *testing.T) {
	// index < 0
	assert.Error(t, (&DoublyLinkedList{}).Insert(-1, 1))

	// index > size
	assert.Error(t, (&DoublyLinkedList{}).Insert(1, 1))

	// []
	list := DoublyLinkedList{}
	list.Insert(0, 1)

	assert.Equal(t, 1, list.size)
	assert.Equal(t, 1, list.head.value)
	assert.Equal(t, list.head, list.tail)
	assert.Nil(t, list.head.next)

	// [1]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Insert(0, 2)

	assert.Equal(t, 2, list.size)
	assert.Equal(t, 2, list.head.value)
	assert.Equal(t, 1, list.tail.value)
	assert.Nil(t, list.tail.next)

	// [1,2]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)
	list.Insert(1, 3)

	assert.Equal(t, 3, list.size)
	assert.Equal(t, 1, list.head.value)
	assert.Equal(t, 2, list.tail.value)
	assert.Nil(t, list.tail.next)
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, (&DoublyLinkedList{}).IsEmpty())
	assert.False(t, (&DoublyLinkedList{size: 1}).IsEmpty())
}

func TestJoin(t *testing.T) {
	// []
	list := DoublyLinkedList{}

	assert.Equal(t, "", list.Join(","))

	// [foo]
	list = DoublyLinkedList{}
	list.Add("foo")

	assert.Equal(t, "foo", list.Join(","))

	// ["foo,bar"]
	list = DoublyLinkedList{}
	list.Add("foo")
	list.Add("bar")

	assert.Equal(t, "foo,bar", list.Join(","))
}

func TestRemove(t *testing.T) {
	// []
	list := DoublyLinkedList{}
	assert.False(t, list.Remove(""))

	// [1]
	list = DoublyLinkedList{}
	list.Add(1)

	assert.False(t, list.Remove(2))
	assert.Equal(t, 1, list.size)

	assert.True(t, list.Remove(1))
	assert.Equal(t, 0, list.size)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

	// [1,2]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)

	assert.True(t, list.Remove(1))
	assert.Equal(t, 1, list.size)
	assert.Equal(t, 2, list.head.value)
	assert.Equal(t, list.tail, list.head)

	// [1,2]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)

	assert.True(t, list.Remove(2))
	assert.Equal(t, 1, list.size)
	assert.Equal(t, 1, list.head.value)
	assert.Equal(t, list.tail, list.head)

	// [1,2,3]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	assert.True(t, list.Remove(2))
	assert.Equal(t, 2, list.size)
	assert.Equal(t, 1, list.head.value)
	assert.Equal(t, 3, list.tail.value)
	assert.NotEqual(t, list.tail, list.head)
}

func TestSize(t *testing.T) {
	assert.Equal(t, 0, DoublyLinkedList{}.size)
	assert.Equal(t, 2, DoublyLinkedList{size: 2}.size)
}