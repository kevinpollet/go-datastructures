/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package doublylinkedlist

import (
	"testing"

	"github.com/kevinpollet/go-datastructures/list"
	"github.com/stretchr/testify/assert"
)

func TestListTypeAssertion(t *testing.T) {
	var doublyLinkedList interface{} = &DoublyLinkedList{}

	cast, ok := doublyLinkedList.(list.List)

	assert.True(t, ok)
	assert.NotNil(t, cast)
}

func TestAdd(t *testing.T) {
	t.Run("ToEmptyList", func(subT *testing.T) {
		list := DoublyLinkedList{}

		list.Add(1)

		assert.Equal(subT, 1, list.size)
		assert.Equal(subT, 1, list.head.value)
		assert.Equal(subT, list.tail, list.head)
	})

	t.Run("ToListWithOneValue", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		list.Add(2)

		assert.Equal(subT, 2, list.size)
		assert.Equal(subT, 2, list.tail.value)
		assert.Equal(subT, 1, list.head.value)
		assert.Equal(subT, list.tail, list.head.next)
	})

	t.Run("ToListWithTwoValues", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.tail = &elt{value: 2}
		list.head = &elt{value: 1, next: list.tail}
		list.tail.prev = list.head
		list.size = 2

		list.Add(3)

		assert.Equal(subT, 3, list.size)
		assert.Equal(subT, 3, list.tail.value)
		assert.Equal(subT, 1, list.head.value)
		assert.Equal(subT, 2, list.head.next.value)
		assert.Equal(subT, list.tail, list.head.next.next)
	})
}

func TestClear(t *testing.T) {
	t.Run("EmptyList", func(subT *testing.T) {
		list := DoublyLinkedList{}

		list.Clear()

		assert.Nil(subT, list.head)
		assert.Nil(subT, list.tail)
		assert.Equal(subT, 0, list.size)
	})

	t.Run("ListWithOneValue", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		list.Clear()

		assert.Nil(subT, list.head)
		assert.Nil(subT, list.tail)
		assert.Equal(subT, 0, list.size)
	})

	t.Run("ListWithTwoValues", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.tail = &elt{value: 2}
		list.head = &elt{value: 1, next: list.tail}
		list.tail.prev = list.head
		list.size = 2

		list.Clear()

		assert.Nil(subT, list.head)
		assert.Nil(subT, list.tail)
		assert.Equal(subT, 0, list.size)
	})
}

func TestGet(t *testing.T) {
	t.Run("WithIndexInfZero", func(subT *testing.T) {
		list := DoublyLinkedList{}

		value, err := list.Get(-1)

		assert.Error(subT, err)
		assert.Nil(subT, value)
	})

	t.Run("WithIndexGreaterThanOrEqualToSize", func(subT *testing.T) {
		list := DoublyLinkedList{}

		value, err := list.Get(0)

		assert.Error(subT, err)
		assert.Nil(subT, value)

		value, err = list.Get(1)

		assert.Error(subT, err)
		assert.Nil(subT, value)
	})

	t.Run("WithHeadIndex", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.tail = &elt{value: 2}
		list.head = &elt{value: 1, next: list.tail}
		list.tail.prev = list.head
		list.size = 2

		value, err := list.Get(0)

		assert.NoError(subT, err)
		assert.Equal(subT, list.head.value, value)
	})

	t.Run("WithTailIndex", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.tail = &elt{value: 2}
		list.head = &elt{value: 1, next: list.tail}
		list.tail.prev = list.head
		list.size = 2

		value, err := list.Get(1)

		assert.NoError(subT, err)
		assert.Equal(subT, list.tail.value, value)
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("ExistingValue", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		assert.Equal(subT, 0, list.IndexOf(1))
	})

	t.Run("MissingValue", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		assert.Equal(subT, -1, list.IndexOf(2))
	})
}

func TestInsert(t *testing.T) {
	t.Run("WithIndexInfZero", func(subT *testing.T) {
		list := DoublyLinkedList{}

		assert.Error(subT, list.Insert(-1, 1))
		assert.Equal(subT, 0, list.size)
	})

	t.Run("WithIndexGreaterThanSize", func(subT *testing.T) {
		list := DoublyLinkedList{}

		assert.Error(subT, list.Insert(-1, 1))
		assert.Equal(subT, 0, list.size)
	})

	t.Run("InEmptyList", func(subT *testing.T) {
		list := DoublyLinkedList{}

		err := list.Insert(0, 1)

		assert.NoError(subT, err)
		assert.Equal(subT, 1, list.size)
		assert.Equal(subT, 1, list.head.value)
		assert.Equal(subT, list.tail, list.head)
		assert.Nil(subT, list.head.prev)
		assert.Nil(subT, list.head.next)
		assert.Nil(subT, list.tail.next)
		assert.Nil(subT, list.tail.prev)
	})

	t.Run("ToHead", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		err := list.Insert(0, 2)

		assert.NoError(subT, err)
		assert.Equal(subT, 2, list.size)
		assert.Equal(subT, 2, list.head.value)
		assert.Equal(subT, 1, list.tail.value)
		assert.Nil(subT, list.head.prev)
		assert.Nil(subT, list.tail.next)
		assert.Equal(subT, list.tail, list.head.next)
		assert.Equal(subT, list.head, list.tail.prev)
	})

	t.Run("ToTail", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		err := list.Insert(1, 2)

		assert.NoError(subT, err)
		assert.Equal(subT, 2, list.size)
		assert.Equal(subT, 1, list.head.value)
		assert.Equal(subT, 2, list.tail.value)
		assert.Nil(subT, list.head.prev)
		assert.Nil(subT, list.tail.next)
		assert.Equal(subT, list.tail, list.head.next)
		assert.Equal(subT, list.head, list.tail.prev)
	})

	t.Run("InListWithTwoValues", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.tail = &elt{value: 2}
		list.head = &elt{value: 1, next: list.tail}
		list.tail.prev = list.head
		list.size = 2

		err := list.Insert(1, 3)

		assert.NoError(subT, err)
		assert.Equal(subT, 3, list.size)
		assert.Equal(subT, 1, list.head.value)
		assert.Equal(subT, 3, list.head.next.value)
		assert.Equal(subT, 2, list.tail.value)
		assert.Nil(subT, list.head.prev)
		assert.Nil(subT, list.tail.next)
		assert.Equal(subT, 3, list.head.next.value)
		assert.Equal(subT, list.head.next, list.tail.prev)
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("WithEmptyList", func(subT *testing.T) {
		list := DoublyLinkedList{}

		assert.True(subT, list.IsEmpty())
	})

	t.Run("WithNonEmptyList", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		assert.False(subT, list.IsEmpty())
	})
}

func TestRemove(t *testing.T) {
	t.Run("WithIndexInfZero", func(subT *testing.T) {
		list := DoublyLinkedList{}

		value, err := list.Remove(-1)

		assert.Error(subT, err)
		assert.Nil(subT, value)
	})

	t.Run("WithIndexGreaterThanOrEqualToSize", func(subT *testing.T) {
		list := DoublyLinkedList{}

		value, err := list.Remove(0)

		assert.Error(subT, err)
		assert.Nil(subT, value)

		value, err = list.Remove(1)

		assert.Error(subT, err)
		assert.Nil(subT, value)
	})

	t.Run("Head", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		value, err := list.Remove(0)

		assert.NoError(subT, err)
		assert.Equal(subT, 1, value)
		assert.Equal(subT, 0, list.size)
		assert.Nil(subT, list.head)
		assert.Nil(subT, list.tail)
	})

	t.Run("HeadInListWithTwoValues", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.tail = &elt{value: 2}
		list.head = &elt{value: 1, next: list.tail}
		list.tail.prev = list.head
		list.size = 2

		value, err := list.Remove(0)

		assert.NoError(subT, err)
		assert.Equal(subT, 1, value)
		assert.Equal(subT, 1, list.size)
		assert.Equal(subT, 2, list.head.value)
		assert.Equal(subT, list.tail, list.head)
		assert.Nil(subT, list.head.prev)
		assert.Nil(subT, list.tail.next)
	})

	t.Run("TailInListWithTwoValues", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.tail = &elt{value: 2}
		list.head = &elt{value: 1, next: list.tail}
		list.tail.prev = list.head
		list.size = 2

		value, err := list.Remove(1)

		assert.NoError(subT, err)
		assert.Equal(subT, 2, value)
		assert.Equal(subT, 1, list.size)
		assert.Equal(subT, 1, list.head.value)
		assert.Equal(subT, list.tail, list.head)
		assert.Nil(subT, list.head.prev)
		assert.Nil(subT, list.tail.next)
	})

	t.Run("MiddleInListWithThreeValues", func(subT *testing.T) {
		middle := &elt{value: 2}
		list := DoublyLinkedList{}
		list.tail = &elt{value: 3, prev: middle}
		list.head = &elt{value: 1, next: middle}
		middle.next = list.tail
		middle.prev = list.head
		list.size = 3

		value, err := list.Remove(1)

		assert.NoError(subT, err)
		assert.Equal(subT, 2, value)
		assert.Equal(subT, 2, list.size)
		assert.Equal(subT, 1, list.head.value)
		assert.Equal(subT, 3, list.tail.value)
		assert.Nil(subT, list.head.prev)
		assert.Nil(subT, list.tail.next)
		assert.Equal(subT, list.tail, list.head.next)
		assert.Equal(subT, list.head, list.tail.prev)
	})
}

func TestSize(t *testing.T) {
	t.Run("WithEmptyList", func(subT *testing.T) {
		list := DoublyLinkedList{}

		assert.Equal(subT, 0, list.Size())
	})

	t.Run("WithNonEmptyList", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		assert.Equal(subT, 1, list.Size())
	})
}

func TestString(t *testing.T) {
	t.Run("WithEmptyList", func(subT *testing.T) {
		list := DoublyLinkedList{}

		assert.Equal(subT, "[]", list.String())
	})

	t.Run("WithListWithOneValue", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1}
		list.tail = list.head
		list.size = 1

		assert.Equal(subT, "[1]", list.String())
	})

	t.Run("WithListWithTwoValues", func(subT *testing.T) {
		list := DoublyLinkedList{}
		list.head = &elt{value: 1, next: &elt{value: 2}}
		list.tail = list.head.next
		list.size = 2

		assert.Equal(subT, "[1,2]", list.String())
	})
}
