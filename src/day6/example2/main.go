package main

import (
	"fmt"
)

type Carer interface {
	GetName() string
	Run()
	Didi()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BMW) Didi() {
	fmt.Printf("%s is didi\n", p.Name)
}

func main() {
	var a interface{}
	var b int
	var c float32
	a = b
	a = c

	fmt.Printf("type of a %T\n", a)

	var car Carer
	var bmw BMW
	bmw.Name = "BMW"
	car = &bmw
	car.Run()
}
