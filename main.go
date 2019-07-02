package main

import (
	"fmt"

	"github.com/kevinpollet/go-datastructures/pkg/list"
)

func main() {
	list := list.LinkedList{}
	list.Append(3).Append(4)

	fmt.Println(list)
}
