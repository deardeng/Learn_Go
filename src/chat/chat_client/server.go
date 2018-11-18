package main

import (
	"chat/proto"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func processServerMessage(conn net.Conn) {
	for {
		msg, err := readPackage(conn)
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(0)
		}

		var userStatus proto.UserStatusNotify
		json.Unmarshal([]byte(msg.Data), &userStatus)
		if err != nil {
			fmt.Println("unmarshal failed, err:", err)
			return
		}

		switch msg.Cmd {
		case proto.UserStatusNotifyRes:
			updateUserStatus(userStatus)
		}
	}
}
