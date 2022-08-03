package main

import (
	"fmt"

	"github.com/go-general/collect/sets"
)

func Sets() {
	set := sets.NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	fmt.Println(set.Values())

	set.Range(func(obj int) bool {
		fmt.Println(obj)
		return true
	})
}
