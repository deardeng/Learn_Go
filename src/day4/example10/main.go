package main

import (
	"fmt"
	"sort"
)

func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	var keys []int

	for k, _ := range a {
		//fmt.Println(k, v)
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		fmt.Println(v, a[v])
	}

}

func testMapSort1() {
	var a map[string]int
	var b map[int]string
	a = make(map[string]int, 5)
	b = make(map[int]string, 5)

	a["abc"] = 10
	a["efg"] = 11

	for k, v := range a {
		b[v] = k
	}

	fmt.Println(b)
}

func main() {
	testMapSort()
	testMapSort1()
}
