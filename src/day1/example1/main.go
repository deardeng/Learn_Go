package main

import "fmt"

func add(a int, arg ...int) int {
	var sum = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

func main() {
	sum := add(10, 20, 40)
	fmt.Println(sum)
}
