package main

import (
	"day7/example1/balance"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// insts := make([]*balance.Instance, 10)
	var insts []*balance.Instance

	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	// var balancer balance.Balancer
	// var conf = "random"
	// if len(os.Args) > 1 {
	// conf = os.Args[1]
	// }

	// if conf == "random" {
	// 	balancer = &balance.RandomBalance{}
	// 	fmt.Println("use random balancer")
	// } else if conf == "roundrobin" {
	// 	balancer = &balance.RoundRobinBalance{}
	// 	fmt.Println("use roundrobin balancer")
	// }

	// balancer := &balance.RandomBalance{}
	// balancer := &balance.RoundRobinBalance{}

	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
