package main

import (
	"fmt"

	"github.com/go-general/collect/maps"
)

func Maps() {
	m := maps.NewHashMap[int, int]()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)

	fmt.Println(m.Keys())
	fmt.Println(m.Values())

	m.Range(func(k int, v int) bool {
		fmt.Printf("key: %d, value: %d\n", k, v)
		return true
	})
}
