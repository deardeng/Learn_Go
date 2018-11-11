package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1024; i++ {
		go func() {
			for {
				select {
				case <-time.After(time.Microsecond):
					fmt.Println("after")
				}
			}
		}()
	}
	time.Sleep(time.Second * 100)
}
