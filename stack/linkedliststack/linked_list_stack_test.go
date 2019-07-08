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
	stackPackage "github.com/kevinpollet/go-datastructures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStackImplementation(test *testing.T) {
	var stack interface{} = &LinkedListStack{}

	cast, ok := stack.(stackPackage.Stack)

	assert.True(test, ok)
	assert.NotNil(test, cast)
}

func TestClear(test *testing.T) {
	test.Run("should removes all values if the stack is empty", func(subTest *testing.T) {
		stack := LinkedListStack{}
		stack.Clear()

		assert.True(subTest, stack.list.IsEmpty())
	})

	test.Run("should removes all values if the stack is not empty", func(subTest *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}
		stack.Clear()

		assert.True(subTest, stack.list.IsEmpty())
	})
}

func TestIsEmpty(test *testing.T) {
	test.Run("should return true if the stack is empty", func(subTest *testing.T) {
		stack := LinkedListStack{}
		assert.True(subTest, stack.IsEmpty())
	})

	test.Run("should return false if the stack is not empty", func(subTest *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}

		assert.False(subTest, stack.IsEmpty())
	})
}

func TestPeek(test *testing.T) {
	test.Run("should return an error if the stack is empty", func(subTest *testing.T) {
		stack := LinkedListStack{}
		value, err := stack.Peek()

		assert.Nil(subTest, value)
		assert.Error(subTest, err)
	})

	test.Run("should return the value at the top of the stack", func(subTest *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}
		value, err := stack.Peek()

		assert.Equal(subTest, 1, value)
		assert.Nil(subTest, err)
	})
}

func TestPop(test *testing.T) {
	test.Run("should return an error if the stack is empty", func(subTest *testing.T) {
		stack := LinkedListStack{}
		_, err := stack.Pop()

		assert.Error(subTest, err)
	})

	test.Run("should return and remove the element at the top of the stack", func(subTest *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}
		value, err := stack.Pop()

		assert.Equal(subTest, 1, value)
		assert.Nil(subTest, err)
	})
}

func TestPush(test *testing.T) {
	test.Run("should add the element to the top of the stack if the stack is empty", func(subTest *testing.T) {
		stack := LinkedListStack{}
		stack.Push(1)

		value, _ := stack.list.Get(0)
		assert.Equal(subTest, 1, value)
	})

	test.Run("should add the element to the top of the stack, if the stack is not empty", func(subTest *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}
		stack.Push(2)

		value, _ := stack.list.Get(0)
		assert.Equal(subTest, 2, value)
	})
}

func TestSize(test *testing.T) {
	test.Run("should returns 0 if the stack is empty", func(subTest *testing.T) {
		stack := LinkedListStack{}

		assert.Equal(subTest, 0, stack.Size())
	})

	test.Run("should returns the number of values if the stack is not empty", func(subTest *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)

		stack := LinkedListStack{list: list}

		assert.Equal(subTest, 1, stack.Size())
	})
}

func TestString(test *testing.T) {
	test.Run("should returns [] if the stack is empty", func(subTest *testing.T) {
		stack := LinkedListStack{}

		assert.Equal(subTest, "[]", stack.String())
	})

	test.Run("should returns the array of values if the stack is not empty", func(subTest *testing.T) {
		list := singlylinkedlist.SinglyLinkedList{}
		list.Add(1)
		list.Add(2)
		list.Add(3)

		stack := LinkedListStack{list: list}

		assert.Equal(subTest, "[1,2,3]", stack.String())
	})
}
