package main

import (
	"fmt"
)

type Test interface {
	Print()
	Sleep()
}

type Test2 interface {
	Print1()
	Sleep1()
}

type Person struct {
	name string
	age  int
}

func (person *Person) Print() {
	fmt.Println("name:", person.name)
	fmt.Println("age:", person.age)
}

func (person *Person) Print1() {
	fmt.Println("name:", person.name)
	fmt.Println("age:", person.age)
}

func (person *Person) Sleep() {
	fmt.Println("person sleep")
}

func (person *Person) Sleep1() {
	fmt.Println("person sleep")
}

type Student struct {
	name  string
	age   int
	score int
}

func (p *Student) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("score:", p.score)
}

func (p *Student) Sleep() {
	fmt.Println("student sleep")
}

func main() {
	var a Test
	// a.Print()
	fmt.Println(a)
	stu := Student{
		name:  "xxx",
		age:   20,
		score: 88,
	}
	a = &stu
	a.Print()
	a.Sleep()

	per := Person{
		name: "hh",
		age:  100,
	}
	a = &per
	a.Print()
	a.Sleep()
	fmt.Println(a)

	var b Test2
	b = &per
	b.Print1()

}
