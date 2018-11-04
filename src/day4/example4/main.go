package main

import (
	"errors"
	"fmt"
	"time"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := initConfig()
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	for {
		test()
		time.Sleep(time.Second)
	}
	var a []int
	a = append(a, 10, 20, 30)
	a = append(a, a...)
	fmt.Println(a)
}
