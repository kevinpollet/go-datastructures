/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package linkedliststack

import (
	"github.com/kevinpollet/go-datastructures/errors"
	"github.com/kevinpollet/go-datastructures/list/singlylinkedlist"
)

// LinkedListStack implements the Stack ADT with a SinglyLinkedList.
type LinkedListStack struct {
	list singlylinkedlist.SinglyLinkedList
}

// Clear removes all values from the stack.
// Complexity: O(1).
func (stack *LinkedListStack) Clear() {
	stack.list.Clear()
}

// IsEmpty returns true if the stack is empty, false otherwise.
// Complexity: O(1).
func (stack *LinkedListStack) IsEmpty() bool {
	return stack.list.IsEmpty()
}

// Peek returns the value at the top of the stack or an error if the stack is empty.
// Complexity: O(1).
func (stack *LinkedListStack) Peek() (interface{}, error) {
	value, err := stack.list.Get(0)
	if _, ok := err.(*errors.IndexOutOfBoundsError); ok {
		return nil, errors.NewNoSuchElementError("cannot peek a value from an empty stack")
	}

	return value, nil
}

// Pop removes and returns the value at the top of the stack or an error if the stack is empty.
// Complexity: O(1).
func (stack *LinkedListStack) Pop() (interface{}, error) {
	value, err := stack.list.Remove(0)
	if _, ok := err.(*errors.IndexOutOfBoundsError); ok {
		return nil, errors.NewNoSuchElementError("cannot pop a value from an empty stack")
	}

	return value, nil
}

// Push adds the given value to the top of the stack.
// Complexity: O(1).
func (stack *LinkedListStack) Push(value interface{}) {
	stack.list.Insert(0, value)
}

// Size returns the number of values in the stack.
// Complexity: O(1).
func (stack *LinkedListStack) Size() int {
	return stack.list.Size()
}

// String returns a string representation of the stack.
// Complexity: O(1).
func (stack LinkedListStack) String() string {
	return stack.list.String()
}
