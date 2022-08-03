package main

import (
	"fmt"

	"github.com/go-general/collect/lists"
)

func Lists() {
	list := lists.NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	fmt.Println(list.Values())

	list.Range(func(index int, val int) bool {
		fmt.Printf("index: %d, value: %d\n", index, val)
		return true
	})
}
