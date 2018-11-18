package main

import (
	"chat/proto"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func sendTextMessage(conn net.Conn, text string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserSendMessageCmd

	var sendReq proto.UserSendMessageReq
	sendReq.Data = text

	data, err := json.Marshal(sendReq)
	if err != nil {
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data failed")
		return
	}
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
	return
}

func enterTalk(conn net.Conn) {
	// var destUserId int
	var msg string
	fmt.Println("please input text")
	fmt.Scanf("%s\n", &msg)
	sendTextMessage(conn, msg)

}

func enterMenu(conn net.Conn) {
	fmt.Println("1. list online user")
	fmt.Println("2. talk")
	fmt.Println("3. exit")

	var sel int
	fmt.Scanf("%d\n", &sel)
	switch sel {
	case 1:
		outputUserOnline()
	case 2:
		enterTalk(conn)
	case 3:
		return
	}

}