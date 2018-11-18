package main

import (
	"chat/common"
	"chat/proto"
	"fmt"
)

var onlineUserMap map[int]*common.User = make(map[int]*common.User, 16)

func outputUserOnline() {
	fmt.Println("Online users list:")
	for id, _ := range onlineUserMap {
		fmt.Println("user: ", id)
	}
}

func updateUserStatus(userStatus proto.UserStatusNotify) {
	user, ok := onlineUserMap[userStatus.UserId]
	if !ok {
		user = &common.User{}
		user.UserId = userStatus.UserId
	}
	user.Status = userStatus.Status
	onlineUserMap[user.UserId] = user
	outputUserOnline()
}
