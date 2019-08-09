/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package binarysearchtree

import (
	"fmt"

	"github.com/kevinpollet/go-datastructures/errors"
)

type node struct {
	value       int
	left, right *node
}

// BinarySearchTree data structure implementation.
type BinarySearchTree struct {
	root *node
	size int
}

// Add adds the given value to the tree.
func (tree *BinarySearchTree) Add(value int) {
	tree.root = add(tree.root, value)
	tree.size++
}

// Remove removes the given value from the tree.
func (tree *BinarySearchTree) Remove(value int) (bool, error) {
	if tree.IsEmpty() {
		return false, errors.NewNoSuchValueError("cannot remove a value from an empty tree")
	}

	root, removed := remove(tree.root, value)

	tree.root = root
	if removed {
		tree.size--
	}
	return removed, nil
}

// Contains returns true if the tree contains the given value, false otherwise.
func (tree *BinarySearchTree) Contains(value int) bool {
	return !tree.IsEmpty() && contains(tree.root, value)
}

// IsEmpty returns true if the tree has no nodes, false otherwise.
func (tree *BinarySearchTree) IsEmpty() bool {
	return tree.Size() == 0
}

// Size returns the number of values in the tree.
func (tree *BinarySearchTree) Size() int {
	return tree.size
}

func (tree BinarySearchTree) String() string {
	return stringify(tree.root)
}

func add(currentNode *node, value int) *node {
	if currentNode == nil {
		return &node{value: value}
	}

	if value > currentNode.value {
		currentNode.right = add(currentNode.right, value)

	} else if value < currentNode.value {
		currentNode.left = add(currentNode.left, value)
	}
	return currentNode
}

func contains(currentNode *node, value int) bool {
	found := false
	if currentNode != nil {
		if value == currentNode.value {
			return true
		} else if value > currentNode.value {
			return contains(currentNode.right, value)
		} else {
			return contains(currentNode.left, value)
		}
	}
	return found
}

func remove(currentNode *node, value int) (*node, bool) {
	if currentNode == nil {
		return nil, false

	} else if currentNode.value > value {
		left, removed := remove(currentNode.left, value)

		currentNode.left = left
		return currentNode, removed

	} else if currentNode.value < value {
		right, removed := remove(currentNode.right, value)

		currentNode.right = right
		return currentNode, removed

	} else if currentNode.left == nil && currentNode.right == nil {
		return nil, true

	} else if currentNode.left == nil {
		return currentNode.right, true

	} else if currentNode.right == nil {
		return currentNode.left, true

	} else {
		min, right := digMin(currentNode.right)
		return &node{value: min, left: currentNode.left, right: right}, true
	}
}

func digMin(currentNode *node) (int, *node) {
	if currentNode.left == nil {
		return currentNode.value, nil
	}

	min, left := digMin(currentNode.left)
	currentNode.left = left

	return min, currentNode
}

func stringify(currentNode *node) string {
	if currentNode == nil {
		return "nil"
	}
	return fmt.Sprintf("(%v, %v, %v)", currentNode.value, stringify(currentNode.left), stringify(currentNode.right))
}
