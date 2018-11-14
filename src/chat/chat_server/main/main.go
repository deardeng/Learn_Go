package main

import (
	"time"
)

func main() {
	initRedis("127.0.0.1:6379", 16, 1024, time.Second*300)
	initUserMgr()
	runServer("0.0.0.0:10000")
}
