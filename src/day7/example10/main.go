package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\n", e)
		}
	}()

	badCall()
	fmt.Printf("After bad call\n")
}

func main() {
	fmt.Println("Calling test\n")
	test()
	fmt.Println("Test completed\n")
}
