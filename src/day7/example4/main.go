package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/appleId.txt")
	if err != nil {
		fmt.Println("read file err :", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("read string failed, err :", err)
		return
	}
	fmt.Printf("read str succ, ret :%s\n", string(str))
}
