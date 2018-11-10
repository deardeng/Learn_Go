package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("len of args:%d\n", len(os.Args))
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please input conf path")
	flag.IntVar(&logLevel, "d", 0, "please input log level")
	flag.Parse()
	for i, v := range os.Args {
		fmt.Printf("args[%d] = %s\n", i, v)
	}

	fmt.Println("path:", confPath)
	fmt.Println("log level:", logLevel)
}
