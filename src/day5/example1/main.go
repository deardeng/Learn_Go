package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var stu Student
	stu.Name = "deardeng"
	stu.Age = 28
	stu.Score = 99

	fmt.Println(stu)
	fmt.Printf("%p\n", &stu.Name)
	fmt.Printf("%p\n", &stu.Age)
	fmt.Printf("%p\n", &stu.Score)
}
