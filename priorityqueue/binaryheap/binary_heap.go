/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package binaryheap

import (
	"fmt"

	"github.com/kevinpollet/go-datastructures/errors"
)

type node struct {
	value    interface{}
	priority int
}

// BinaryHeap implements the PriorityQueue ADT.
type BinaryHeap struct {
	tree []*node
}

// Clear removes all values from the heap.
// Complexity: O(1)
func (heap *BinaryHeap) Clear() {
	heap.tree = nil
}

// IsEmpty returns true if the heap is empty, false otherwise.
// Complexity: O(1)
func (heap *BinaryHeap) IsEmpty() bool {
	return heap.Size() == 0
}

// Offer adds the given value to the heap with the given priority.
// Complexity: O(log(n))
func (heap *BinaryHeap) Offer(value interface{}, priority int) {
	heap.tree = append(heap.tree, &node{value: value, priority: priority})

	childIndex := heap.Size() - 1
	parentIndex := (childIndex - 1) / 2
	for parentIndex >= 0 && heap.tree[parentIndex].priority > heap.tree[childIndex].priority {
		temp := heap.tree[childIndex]
		heap.tree[childIndex] = heap.tree[parentIndex]
		heap.tree[parentIndex] = temp

		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
}

// Peek returns the value with the hightest priority in the heap or an error if the heap is empty.
// Complexity: O(1)
func (heap *BinaryHeap) Peek() (interface{}, error) {
	if heap.IsEmpty() {
		return nil, errors.NewNoSuchValueError("cannot peek a value from an empty heap")
	}
	return heap.tree[0].value, nil
}

// Poll returns and removes the value with the hightest priority in the heap or an error if the heap is empty.
// Complexity: O(log(n))
func (heap *BinaryHeap) Poll() (interface{}, error) {
	if heap.IsEmpty() {
		return nil, errors.NewNoSuchValueError("cannot poll a value from an empty heap")
	}

	rootNode := heap.tree[0]

	heap.tree[0] = heap.tree[heap.Size()-1]
	heap.tree = heap.tree[0 : heap.Size()-1]

	parentIndex := 0
	leftChildIndex := 2*parentIndex + 1
	rightChildIndex := 2*parentIndex + 2
	for leftChildIndex <= heap.Size()-1 {
		minChildIndex := leftChildIndex
		if rightChildIndex <= heap.Size()-1 && heap.tree[rightChildIndex].priority < heap.tree[leftChildIndex].priority {
			minChildIndex = rightChildIndex
		}

		if heap.tree[parentIndex].priority < heap.tree[minChildIndex].priority {
			break
		}

		temp := heap.tree[parentIndex]
		heap.tree[parentIndex] = heap.tree[minChildIndex]
		heap.tree[minChildIndex] = temp

		parentIndex = minChildIndex
		leftChildIndex = 2*parentIndex + 1
		rightChildIndex = 2*parentIndex + 2
	}

	return rootNode.value, nil
}

// Size returns the number of values in the heap.
// Complexity: O(1)
func (heap *BinaryHeap) Size() int {
	return len(heap.tree)
}

// String returns a string representation of the tree.
// Complexity: O(n)
func (heap BinaryHeap) String() string {
	str := "["
	for index, node := range heap.tree {
		str += fmt.Sprintf("(value: %v, priority: %d)", node.value, node.priority)
		if index != heap.Size()-1 {
			str += ","
		}
	}

	return str + "]"
}
