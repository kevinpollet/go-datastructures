/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedliststack

import (
	"github.com/kevinpollet/go-datastructures/list/singlylinkedlist"
	stackPackage "github.com/kevinpollet/go-datastructures/stack"
)

// LinkedListStack implementation backed by a singly linked list
type LinkedListStack struct {
	list singlylinkedlist.SinglyLinkedList
}

// Clear removes all values in the stack
func (stack *LinkedListStack) Clear() {
	stack.list.Clear()
}

// IsEmpty returns true is the stack is empty, false otherwise
func (stack *LinkedListStack) IsEmpty() bool {
	return stack.list.IsEmpty()
}

// Peek returns the value at the top of the stack, nil if empty
func (stack *LinkedListStack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, stackPackage.NewEmptyStackError()
	}

	value, _ := stack.list.Get(0)
	return value, nil
}

// Pop removes and returns the value at the top of the stack
func (stack *LinkedListStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, stackPackage.NewEmptyStackError()
	}

	value, _ := stack.list.Get(0)
	stack.list.Remove(value)

	return value, nil
}

// Push adds a value to the top of the stack
func (stack *LinkedListStack) Push(value interface{}) {
	stack.list.Insert(0, value)
}

// Size returns the number of values in the stack
func (stack *LinkedListStack) Size() int {
	return stack.list.Size()
}

func (stack LinkedListStack) String() string {
	return stack.list.String()
}
