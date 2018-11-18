package main

import (
	"fmt"
	"net"
)

var userId int
var passwd string

func main() {
	// var userId int
	// var passwd string
	fmt.Scanf("%d %s\n", &userId, &passwd)
	conn, err := net.Dial("tcp", "localhost:10001")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	err = login(conn, userId, passwd)
	if err != nil {
		fmt.Println("login failed, err: ", err)
		return
	}

	outputUserOnline()
	go processServerMessage(conn)
	for {
		enterMenu(conn)
	}

}
