package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection:keep-alive\r\n"
	msg += "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36"
	msg += "\r\n\r\n"

	n, err := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed, ", err)
		return
	}
	fmt.Println("send to baidu bytes:", n)
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("count:", count, "err: ", err)
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net"
// )

// func main() {

// 	conn, err := net.Dial("tcp", "www.baidu.com:80")
// 	if err != nil {
// 		fmt.Println("Error dialing", err.Error())
// 		return
// 	}
// 	defer conn.Close()
// 	msg := "GET / HTTP/1.1\r\n"
// 	msg += "Host:www.baidu.com\r\n"
// 	msg += "Connection:keep-alive\r\n"
// 	//msg += "User-Agent:Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36\r\n"
// 	msg += "\r\n\r\n"

// 	//io.WriteString(os.Stdout, msg)
// 	n, err := io.WriteString(conn, msg)
// 	if err != nil {
// 		fmt.Println("write string failed, ", err)
// 		return
// 	}
// 	fmt.Println("send to baidu.com bytes:", n)
// 	buf := make([]byte, 4096)
// 	for {
// 		count, err := conn.Read(buf)
// 		fmt.Println("count:", count, "err:", err)
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(buf[0:count]))
// 	}
// }
