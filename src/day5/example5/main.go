package main

import (
	"fmt"
)

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	wheel int
}

type Train struct {
	c Car
}

func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s] weight=[%d]", p.c.name, p.c.weight)
	return str
}

func main() {

	var a Bike
	a.weight = 100
	a.name = "bike"
	a.wheel = 2

	fmt.Println(a)
	a.Run()

	var b Train
	b.c.weight = 1000
	b.c.name = "train"
	// fmt.Println(b)
	b.c.Run()
	fmt.Printf("%s\n", &b)
	fmt.Println(&b)
}
