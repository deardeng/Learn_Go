package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	c := val.Int()
	fmt.Printf("get value interface{} %d\n", c)
}

func main() {
	var a int = 200
	test(a)

	var b Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 98,
	}
	test(b)
	testInt(12212)
}
