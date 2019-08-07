/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package unionfind

import (
	"fmt"
)

// UnionFind data structure implementation.
type UnionFind struct {
	ids   []int
	sizes []int
}

// New constructs and returns a pointer to an UnionFind data structure instance.
func New(size int) *UnionFind {
	unionFind := UnionFind{
		ids:   make([]int, size),
		sizes: make([]int, size),
	}

	for i := range unionFind.ids {
		unionFind.ids[i] = i
		unionFind.sizes[i] = 1
	}

	return &unionFind
}

// IsConnected returns true if the given elements are in the same set, false otherwise.
func (unionFind *UnionFind) IsConnected(p, q int) bool {
	return unionFind.Find(p) == unionFind.Find(q)
}

// Find returns the set of the given element.
func (unionFind *UnionFind) Find(p int) int {
	root := p
	for root != unionFind.ids[root] {
		root = unionFind.ids[root]
	}

	// Compress path leading to the root (Path compression optimization).
	for next := p; next != root; {
		temp := unionFind.ids[next]
		unionFind.ids[next] = root
		next = temp
	}

	return root
}

// Size returns the number of elements in this UnionFind.
func (unionFind *UnionFind) Size() int {
	return len(unionFind.ids)
}

func (unionFind UnionFind) String() string {
	str := "["
	for i, value := range unionFind.ids {
		str += fmt.Sprintf("(value: %d, size: %d)", value, unionFind.sizes[i])
		if i != len(unionFind.ids)-1 {
			str += ","
		}
	}
	return str + "]"
}

// Unify connects the given elements' roots.
// The smaller set is merged into the bigger one.
func (unionFind *UnionFind) Unify(p, q int) {
	pRoot := unionFind.Find(p)
	qRoot := unionFind.Find(q)

	if pRoot != qRoot {
		if unionFind.sizes[pRoot] > unionFind.sizes[qRoot] {
			unionFind.ids[qRoot] = pRoot
			unionFind.sizes[pRoot] += unionFind.sizes[qRoot]
		} else {
			unionFind.ids[pRoot] = qRoot
			unionFind.sizes[qRoot] += unionFind.sizes[pRoot]
		}
	}
}
