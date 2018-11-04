package main

import (
	"fmt"
	"strings"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func test(arr *[5]int) {
	(*arr)[0] = 1000
}

func testArray() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{3, 8, 3, 4, 5, 0}
	var a3 = [...]int{1: 100, 4: 20039}

	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

func testArray2() {
	var a [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	for row, v := range a {
		for col, v1 := range v {
			fmt.Printf("(%d, %d)=%d ", row, col, v1)
		}
		fmt.Println()
	}
}

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))

	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("abc.bmp"))

	var a [5]int
	test(&a)
	fmt.Println(a)

	testArray()

	testArray2()
}
