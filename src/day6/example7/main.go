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
	val.Elem().SetInt(100)
	c := val.Elem().Int()
	fmt.Printf("get value interface{} %d\n", c)
	fmt.Printf("string val : %s\n", val.Elem().String())
	fmt.Printf("string val : %d\n", val.Elem().Int())
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

	var c int = 1
	testInt(&c)
	fmt.Println(c)
}
