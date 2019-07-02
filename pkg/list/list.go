/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package list

import (
	"fmt"
)

type cell struct {
	next  *cell
	value interface{}
}

type LinkedList struct {
	head *cell
	tail *cell
	size int
}

func (l *LinkedList) Append(value interface{}) {
	return
}

func (l LinkedList) Join(delim string) (result string) {
	result = ""
	for it := l.head; it != nil; {
		result += fmt.Sprintf("%s,", it.value)
		it = it.next
	}
	return
}

func (l LinkedList) Size() int {
	return l.size
}

func (l LinkedList) String() string {
	return fmt.Sprintf("[%s]")
}
