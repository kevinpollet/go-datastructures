/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package doublylinkedlist

import (
	"testing"

	listPackage "github.com/kevinpollet/go-datastructures/list"
	"github.com/stretchr/testify/assert"
)

func TestListImplementation(test *testing.T) {
	var list interface{} = &DoublyLinkedList{}

	cast, ok := list.(listPackage.List)

	assert.True(test, ok)
	assert.NotNil(test, cast)
}

func TestAdd(test *testing.T) {
	// []
	list := DoublyLinkedList{}
	list.Add(1)

	assert.Equal(test, 1, list.size)
	assert.Equal(test, 1, list.head.value)
	assert.Equal(test, list.tail, list.head)
	assert.Nil(test, list.tail.next)

	// [1]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)

	assert.Equal(test, 2, list.size)
	assert.Equal(test, 1, list.head.value)
	assert.Equal(test, list.tail, list.head.next)
	assert.Equal(test, 2, list.tail.value)
	assert.Nil(test, list.tail.next)
}

func TestAddAtIndex(test *testing.T) {
	// index < 0
	assert.Error(test, (&DoublyLinkedList{}).AddAtIndex(-1, 1))

	// index > size
	assert.Error(test, (&DoublyLinkedList{}).AddAtIndex(1, 1))

	// []
	list := DoublyLinkedList{}
	list.AddAtIndex(0, 1)

	assert.Equal(test, 1, list.size)
	assert.Equal(test, 1, list.head.value)
	assert.Equal(test, list.head, list.tail)
	assert.Nil(test, list.head.next)

	// [1]
	list = DoublyLinkedList{}
	list.Add(1)
	list.AddAtIndex(0, 2)

	assert.Equal(test, 2, list.size)
	assert.Equal(test, 2, list.head.value)
	assert.Equal(test, 1, list.tail.value)
	assert.Nil(test, list.tail.next)

	// [1,2]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)
	list.AddAtIndex(1, 3)

	assert.Equal(test, 3, list.size)
	assert.Equal(test, 1, list.head.value)
	assert.Equal(test, 2, list.tail.value)
	assert.Nil(test, list.tail.next)
}

func TestClear(test *testing.T) {
	// []
	list := DoublyLinkedList{}
	list.Clear()

	assert.Equal(test, 0, list.size)
	assert.Nil(test, list.head)
	assert.Nil(test, list.tail)

	// [1]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Clear()

	assert.Equal(test, 0, list.size)
	assert.Nil(test, list.head)
	assert.Nil(test, list.tail)

	// [1, 2]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)
	list.Clear()

	assert.Equal(test, 0, list.size)
	assert.Nil(test, list.head)
	assert.Nil(test, list.tail)
}

func TestContains(test *testing.T) {
	// []
	assert.False(test, (&DoublyLinkedList{}).Contains("foo"))

	// ["foo"]
	list := DoublyLinkedList{}
	list.Add("foo")

	assert.True(test, list.Contains("foo"))
	assert.False(test, list.Contains("bar"))
}

func TestGet(test *testing.T) {
	// index < 0
	_, err := (&DoublyLinkedList{}).Get(-1)

	assert.Error(test, err)

	// index >= size
	_, err = (&DoublyLinkedList{size: 0}).Get(0)

	assert.Error(test, err, "Index Out of Bounds Error")

	// [1]
	list := DoublyLinkedList{}
	list.Add(1)
	res, err := list.Get(0)

	assert.Equal(test, 1, res)
	assert.Nil(test, err)

}

func TestIsEmpty(test *testing.T) {
	assert.True(test, (&DoublyLinkedList{}).IsEmpty())
	assert.False(test, (&DoublyLinkedList{size: 1}).IsEmpty())
}

func TestJoin(test *testing.T) {
	// []
	list := DoublyLinkedList{}

	assert.Equal(test, "", list.Join(","))

	// [foo]
	list = DoublyLinkedList{}
	list.Add("foo")

	assert.Equal(test, "foo", list.Join(","))

	// ["foo,bar"]
	list = DoublyLinkedList{}
	list.Add("foo")
	list.Add("bar")

	assert.Equal(test, "foo,bar", list.Join(","))
}

func TestRemove(test *testing.T) {
	// []
	list := DoublyLinkedList{}
	assert.False(test, list.Remove(""))

	// [1]
	list = DoublyLinkedList{}
	list.Add(1)

	assert.False(test, list.Remove(2))
	assert.Equal(test, 1, list.size)

	assert.True(test, list.Remove(1))
	assert.Equal(test, 0, list.size)
	assert.Nil(test, list.head)
	assert.Nil(test, list.tail)

	// [1,2]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)

	assert.True(test, list.Remove(1))
	assert.Equal(test, 1, list.size)
	assert.Equal(test, 2, list.head.value)
	assert.Equal(test, list.tail, list.head)

	// [1,2]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)

	assert.True(test, list.Remove(2))
	assert.Equal(test, 1, list.size)
	assert.Equal(test, 1, list.head.value)
	assert.Equal(test, list.tail, list.head)

	// [1,2,3]
	list = DoublyLinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	assert.True(test, list.Remove(2))
	assert.Equal(test, 2, list.size)
	assert.Equal(test, 1, list.head.value)
	assert.Equal(test, 3, list.tail.value)
	assert.NotEqual(test, list.tail, list.head)
}

func TestSize(test *testing.T) {
	assert.Equal(test, 0, DoublyLinkedList{}.size)
	assert.Equal(test, 2, DoublyLinkedList{size: 2}.size)
}
