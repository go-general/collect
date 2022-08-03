# collect
common collect util



## List

```go
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
```



## Map

```go
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
```



## Set

```go
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
```
