package main

import (
	"fmt"

	"github.com/go-general/collect/maps"
)

func Maps() {
	m := maps.NewHashMap[string, int64]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	m.Range(func(k string, v int64) bool {
		fmt.Println(k, v)
		return true
	})

	m.Remove("b")

	m.Range(func(k string, v int64) bool {
		fmt.Println(k, v)
		return true
	})

	fmt.Println(m.Values())
}
