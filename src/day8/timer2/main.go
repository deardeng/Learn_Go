package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				t := time.NewTicker(time.Second)
				select {
				case <-t.C:
					fmt.Println("timeout")
				}
				t.Stop()
			}
		}()
	}
	time.Sleep(time.Second * 100)
}
