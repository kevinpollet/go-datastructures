/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package list

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	list := LinkedList{}

	if !list.IsEmpty() {
		t.Errorf("List should be empty")
	}

	list.Add(1)
	if list.IsEmpty() {
		t.Errorf("List should not be empty")
	}
}
