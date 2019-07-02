package main

import (
	"fmt"

	"github.com/kevinpollet/go-datastructures/pkg/list"
)

func main() {
	list := (&list.LinkedList{}).
		Add(1).
		Add(2).
		Add(3).
		//Clear().
		Add(4).
		Add(5)

	list.Insert(4, 10)
	fmt.Println(list)

	// fmt.Println(list.Remove(5))
	fmt.Println(list.Size())
	fmt.Println(list.Contains(1))
	fmt.Println(list.Join("|"))
	fmt.Println(list.Get(0))
}
