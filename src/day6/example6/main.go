package main

import "fmt"

func main() {
	// var intLink Link
	var sLink Link
	for i := 0; i < 10; i++ {
		// intLink.InsertHead(i)
		// intLink.InsertTail(i)
		sLink.InsertHead(fmt.Sprintf("str %d", i))
	}
	// intLink.Trans()
	sLink.Trans()
}
