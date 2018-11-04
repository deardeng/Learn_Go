package main

import (
	"fmt"
)

type slice struct {
	ptr *[1000]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new([1000]int)
	s.cap = cap
	s.len = 0
	return s
}

func testSlice() {
	var slice []int
	arr := [...]int{1, 2, 3, 4, 5}

	slice = arr[:]
	slice = slice[1:]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func modify(s slice) {
	s.ptr[1] = 10000
}

func testSlice2() {
	var s1 slice
	s1 = make1(s1, 10)

	s1.ptr[0] = 100
	modify(s1)
	fmt.Println(s1.ptr)
}

func modify1(a []int) {
	a[1] = 100
}

func testSlice3() {
	b := []int{1, 2, 3, 4}
	modify1(b)
	fmt.Println(b)
}

func testSlice4() {
	var a = [10]int{1, 2, 3, 4, 5}
	b := a[1:5]
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a[1])
}

func testSlice5() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]

	s[1] = 100
	fmt.Println("before a: ", a)

	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])

	s[1] = 1000
	fmt.Println("after a: ", a)

	var c = []int{1, 2, 3}
	var b = []int{4, 5, 6}
	c = append(c, b...)
	fmt.Println(c)
}

func testCopy() {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 1)
	copy(b, a)

	fmt.Println(b)
}

func testString() {
	s := "hello world"
	s1 := s[0:5]
	s2 := s[6:]

	fmt.Println(s1)
	fmt.Println(s2)
}

func testModifyString() {
	s := "hello world"
	s1 := []rune(s)

	s1[1] = '0'
	str := string(s1)
	fmt.Println(str)
}

func main() {
	testSlice()
	testSlice2()
	testSlice3()
	testSlice4()
	testSlice5()
	testCopy()
	testString()

	testModifyString()
}
