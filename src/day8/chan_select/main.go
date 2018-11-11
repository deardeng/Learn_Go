package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)

	ch2 := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i * i
		}
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		// default:
		case <-time.After(3 * time.Second):
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}
