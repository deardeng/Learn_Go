package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1, 2, 98988, 8, 212, 440, 21212}
	sort.Ints(a[:])

	fmt.Println(a)
}

func testStrings() {
	var a = [...]string{"abc", "efg", "b", "A", "eeee"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2.3, 0.8, 28.1, 2932.1, 220.1}
	sort.Float64s(a[:])

	fmt.Println(a)
}

func testIntSearch() {
	var a = [...]int{1, 8, 10, 22, 44, 828}

	index := sort.SearchInts(a[:], 1)
	fmt.Println(index)
}

func main() {
	testIntSort()
	testStrings()
	testFloat()
	testIntSearch()
}
