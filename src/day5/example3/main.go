package main

import (
	"fmt"
	"time"
)

type Cart struct {
	name string
	age  int
}

type Train struct {
	Cart
	int
	start time.Time
}

func main() {

	var t Train
	t.name = "train"
	t.age = 100
	t.int = 200
	t.start = time.Now()

	fmt.Println(t)
}
