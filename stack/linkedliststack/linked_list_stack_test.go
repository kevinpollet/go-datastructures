/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedliststack

import (
	"testing"

	"github.com/kevinpollet/go-datastructures/list/singlylinkedlist"
	"github.com/kevinpollet/go-datastructures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStackTypeAssertion(t *testing.T) {
	var linkedListStack interface{} = &LinkedListStack{}

	cast, ok := linkedListStack.(stack.Stack)

	assert.True(t, ok)
	assert.NotNil(t, cast)
}

func TestClear(t *testing.T) {
	t.Run("EmptyStack", func(subT *testing.T) {
		stack := LinkedListStack{}
		stack.Clear()

		assert.True(subT, stack.list.IsEmpty())
	})

	t.Run("NonEmptyStack", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}
		stack.Clear()

		assert.True(subT, stack.list.IsEmpty())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("EmptyStack", func(subT *testing.T) {
		stack := LinkedListStack{}

		assert.True(subT, stack.IsEmpty())
	})

	t.Run("NonEmptyStack", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}

		assert.False(subT, stack.IsEmpty())
	})
}

func TestPeek(t *testing.T) {
	t.Run("EmptyStack", func(subT *testing.T) {
		stack := LinkedListStack{}

		value, err := stack.Peek()

		assert.Nil(subT, value)
		assert.Error(subT, err)
	})

	t.Run("StackWithOneValue", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}
		value, err := stack.Peek()

		assert.Equal(subT, 1, value)
		assert.Equal(subT, 1, stack.list.Size())
		assert.NoError(subT, err)
	})
}

func TestPop(t *testing.T) {
	t.Run("EmptyStack", func(subT *testing.T) {
		stack := LinkedListStack{}

		value, err := stack.Pop()

		assert.Error(subT, err)
		assert.Nil(subT, value)
	})

	t.Run("StackWithOneValue", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}
		value, err := stack.Pop()

		assert.Equal(subT, 1, value)
		assert.Equal(subT, 0, stack.list.Size())
		assert.NoError(subT, err)
	})

	t.Run("StackWithTwoValues", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)
		list.Add(2)

		stack := LinkedListStack{list: list}
		value, err := stack.Pop()

		assert.Equal(subT, 1, value)
		assert.Equal(subT, 1, stack.list.Size())
		assert.NoError(subT, err)
	})
}

func TestPush(t *testing.T) {
	t.Run("EmptyStack", func(subT *testing.T) {
		stack := LinkedListStack{}
		stack.Push(1)

		value, _ := stack.list.Get(0)
		assert.Equal(subT, 1, value)
	})

	t.Run("StackWithOneValue", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}
		stack.Push(2)

		value, _ := stack.list.Get(0)
		assert.Equal(subT, 2, value)
	})
}

func TestSize(t *testing.T) {
	t.Run("EmptyStack", func(subT *testing.T) {
		stack := LinkedListStack{}

		assert.Equal(subT, 0, stack.Size())
	})

	t.Run("StackWithOneValue", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}

		assert.Equal(subT, 1, stack.Size())
	})
}

func TestString(t *testing.T) {
	t.Run("EmptyStack", func(subT *testing.T) {
		stack := LinkedListStack{}

		assert.Equal(subT, "[]", stack.String())
	})

	t.Run("StackWithOneValue", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}

		assert.Equal(subT, "[1]", stack.String())
	})

	t.Run("StackWithTwoValues", func(subT *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)
		list.Add(2)

		stack := LinkedListStack{list: list}

		assert.Equal(subT, "[1,2]", stack.String())
	})
}
