package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic: ", err)
		}
	}()
	var m map[string]int
	m["stu"] = 100
}

func calc() {
	for {
		fmt.Println("i'm calc")
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	for i := 0; i < 100; i++ {
		go calc()
	}
	time.Sleep(time.Second)
}
